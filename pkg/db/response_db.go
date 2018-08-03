package db

import (
	"github.com/james-millner/go-lang-web-app/pkg/model"

	"github.com/jinzhu/gorm"
)

// Response interface for getting responses from MySQL.
type Response interface {
	FindBySourceURLAndURLFound(source string, url string) *model.Response
	FindAll() []*model.Response
	Save(r *model.Response) *model.Response
}

//CaseStudy interface for handling communication with the CaseStudy entity.
type CaseStudy interface {
	FindByID(id uint) *model.CaseStudy
	Save(c *model.CaseStudy) *model.CaseStudy
}

//DB Struct for communication with GORM.
type DB struct {
	db *gorm.DB
}

// New creates a new instance of DB and returns a reference to it
func New(db *gorm.DB) *DB {
	return &DB{
		db: db,
	}
}

//FindBySourceURLAndURLFound method
func (d *DB) FindBySourceURLAndURLFound(sourceurl string, urlfound string) *model.Response {
	var c model.Response
	c.URLFound = urlfound
	c.SourceURL = sourceurl
	d.db.Where(&c).First(&c)

	return &c
}

//FindAll method
func (d *DB) FindAll() []*model.Response {
	var responses []*model.Response
	d.db.Find(&responses)

	return responses
}

//SaveResponse method
func (d *DB) SaveResponse(r *model.Response) *model.Response {
	if d.db.NewRecord(r) {
		d.db.Create(&r)
	} else {
		d.db.Save(&r)
	}

	return r
}

//FindByID method
func (d *DB) FindByID(id string) *model.CaseStudy {
	var c model.CaseStudy
	c.ID = id
	d.db.Where(&c).First(&c)

	return &c
}

//SaveCaseStudy function; to save case study objects
func (d *DB) SaveCaseStudy(c *model.CaseStudy) *model.CaseStudy {
	if d.db.NewRecord(c) {
		d.db.Create(&c)
	} else {
		d.db.Save(&c)
	}

	return c
}

//SaveCaseStudyOrganisation function; to save organisations against a Case Study
func (d *DB) SaveCaseStudyOrganisation(c *model.CaseStudyOrganisations) *model.CaseStudyOrganisations {
	if d.db.NewRecord(c) {
		d.db.Create(&c)
	} else {
		d.db.Save(&c)
	}

	return c
}

//SaveCaseStudyPerson function; to save people against a Case Study
func (d *DB) SaveCaseStudyPerson(c *model.CaseStudyPeople) *model.CaseStudyPeople {
	if d.db.NewRecord(c) {
		d.db.Create(&c)
	} else {
		d.db.Save(&c)
	}

	return c
}

//SaveCaseStudyLocation function; to save locations against a Case Study
func (d *DB) SaveCaseStudyLocation(c *model.CaseStudyLocations) *model.CaseStudyLocations {
	if d.db.NewRecord(c) {
		d.db.Create(&c)
	} else {
		d.db.Save(&c)
	}

	return c
}

//FindCaseStudyBySourceAndCompanyNumber function
func (d *DB) FindCaseStudyBySourceAndCompanyNumber(source string, companyNumber string) *model.CaseStudy {
	var c model.CaseStudy
	c.SourceURL = source
	c.CompanyNumber = companyNumber
	d.db.Where(&c).First(&c)

	return &c

}

//DeleteCaseStudyOrganisations function
func (d *DB) DeleteCaseStudyOrganisations(caseStudyId string) {
	d.db.Where("case_study_id = ?", caseStudyId).Delete(model.CaseStudyOrganisations{})
}

//DeleteCaseStudyPeople function
func (d *DB) DeleteCaseStudyPeople(caseStudyId string) {
	d.db.Where("case_study_id = ?", caseStudyId).Delete(model.CaseStudyPeople{})
}

//DeleteCaseStudyLocation function
func (d *DB) DeleteCaseStudyLocations(caseStudyId string) {
	d.db.Where("case_study_id = ?", caseStudyId).Delete(model.CaseStudyLocations{})
}

//FindCaseStudyOrganisationByNameAndCaseID function
func (d *DB) FindCaseStudyOrganisationByNameAndCaseID(organisationName string, id string) *model.CaseStudyOrganisations {
	var c model.CaseStudyOrganisations
	c.OrganisationName = organisationName
	c.CaseStudyID = id
	d.db.Where(&c).First(&c)

	return &c
}

//FindCaseStudyPersonByNameAndCaseID function
func (d *DB) FindCaseStudyPersonByNameAndCaseID(personName string, id string) *model.CaseStudyPeople {
	var c model.CaseStudyPeople
	c.PersonName = personName
	c.CaseStudyID = id
	d.db.Where(&c).First(&c)

	return &c
}

//FindCaseStudyLocationByLocationAndCaseID function
func (d *DB) FindCaseStudyLocationByLocationAndCaseID(location string, id string) *model.CaseStudyLocations {
	var c model.CaseStudyLocations
	c.Location = location
	c.CaseStudyID = id
	d.db.Where(&c).First(&c)

	return &c
}
