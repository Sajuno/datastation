// Copyright 2016 - 2022 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to and
// read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and
// writing spreadsheet documents generated by Microsoft Excel™ 2007 and later.
// Supports complex components by high compatibility, and provided streaming
// API for generating or reading data from a worksheet with huge amounts of
// data. This library needs Go version 1.15 or later.

package excelize

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
)

// SetAppProps provides a function to set document application properties. The
// properties that can be set are:
//
//     Property          | Description
//    -------------------+--------------------------------------------------------------------------
//     Application       | The name of the application that created this document.
//                       |
//     ScaleCrop         | Indicates the display mode of the document thumbnail. Set this element
//                       | to 'true' to enable scaling of the document thumbnail to the display. Set
//                       | this element to 'false' to enable cropping of the document thumbnail to
//                       | show only sections that will fit the display.
//                       |
//     DocSecurity       | Security level of a document as a numeric value. Document security is
//                       | defined as:
//                       | 1 - Document is password protected.
//                       | 2 - Document is recommended to be opened as read-only.
//                       | 3 - Document is enforced to be opened as read-only.
//                       | 4 - Document is locked for annotation.
//                       |
//     Company           | The name of a company associated with the document.
//                       |
//     LinksUpToDate     | Indicates whether hyperlinks in a document are up-to-date. Set this
//                       | element to 'true' to indicate that hyperlinks are updated. Set this
//                       | element to 'false' to indicate that hyperlinks are outdated.
//                       |
//     HyperlinksChanged | Specifies that one or more hyperlinks in this part were updated
//                       | exclusively in this part by a producer. The next producer to open this
//                       | document shall update the hyperlink relationships with the new
//                       | hyperlinks specified in this part.
//                       |
//     AppVersion        | Specifies the version of the application which produced this document.
//                       | The content of this element shall be of the form XX.YYYY where X and Y
//                       | represent numerical values, or the document shall be considered
//                       | non-conformant.
//
// For example:
//
//    err := f.SetAppProps(&excelize.AppProperties{
//        Application:       "Microsoft Excel",
//        ScaleCrop:         true,
//        DocSecurity:       3,
//        Company:           "Company Name",
//        LinksUpToDate:     true,
//        HyperlinksChanged: true,
//        AppVersion:        "16.0000",
//    })
//
func (f *File) SetAppProps(appProperties *AppProperties) (err error) {
	var (
		app                *xlsxProperties
		fields             []string
		output             []byte
		immutable, mutable reflect.Value
		field              string
	)
	app = new(xlsxProperties)
	if err = f.xmlNewDecoder(bytes.NewReader(namespaceStrictToTransitional(f.readXML(defaultXMLPathDocPropsApp)))).
		Decode(app); err != nil && err != io.EOF {
		err = fmt.Errorf("xml decode error: %s", err)
		return
	}
	fields = []string{"Application", "ScaleCrop", "DocSecurity", "Company", "LinksUpToDate", "HyperlinksChanged", "AppVersion"}
	immutable, mutable = reflect.ValueOf(*appProperties), reflect.ValueOf(app).Elem()
	for _, field = range fields {
		immutableField := immutable.FieldByName(field)
		switch immutableField.Kind() {
		case reflect.Bool:
			mutable.FieldByName(field).SetBool(immutableField.Bool())
		case reflect.Int:
			mutable.FieldByName(field).SetInt(immutableField.Int())
		default:
			mutable.FieldByName(field).SetString(immutableField.String())
		}
	}
	app.Vt = NameSpaceDocumentPropertiesVariantTypes.Value
	output, err = xml.Marshal(app)
	f.saveFileList(defaultXMLPathDocPropsApp, output)
	return
}

// GetAppProps provides a function to get document application properties.
func (f *File) GetAppProps() (ret *AppProperties, err error) {
	app := new(xlsxProperties)
	if err = f.xmlNewDecoder(bytes.NewReader(namespaceStrictToTransitional(f.readXML(defaultXMLPathDocPropsApp)))).
		Decode(app); err != nil && err != io.EOF {
		err = fmt.Errorf("xml decode error: %s", err)
		return
	}
	ret, err = &AppProperties{
		Application:       app.Application,
		ScaleCrop:         app.ScaleCrop,
		DocSecurity:       app.DocSecurity,
		Company:           app.Company,
		LinksUpToDate:     app.LinksUpToDate,
		HyperlinksChanged: app.HyperlinksChanged,
		AppVersion:        app.AppVersion,
	}, nil
	return
}

// SetDocProps provides a function to set document core properties. The
// properties that can be set are:
//
//     Property       | Description
//    ----------------+-----------------------------------------------------------------------------
//     Title          | The name given to the resource.
//                    |
//     Subject        | The topic of the content of the resource.
//                    |
//     Creator        | An entity primarily responsible for making the content of the resource.
//                    |
//     Keywords       | A delimited set of keywords to support searching and indexing. This is
//                    | typically a list of terms that are not available elsewhere in the properties.
//                    |
//     Description    | An explanation of the content of the resource.
//                    |
//     LastModifiedBy | The user who performed the last modification. The identification is
//                    | environment-specific.
//                    |
//     Language       | The language of the intellectual content of the resource.
//                    |
//     Identifier     | An unambiguous reference to the resource within a given context.
//                    |
//     Revision       | The topic of the content of the resource.
//                    |
//     ContentStatus  | The status of the content. For example: Values might include "Draft",
//                    | "Reviewed" and "Final"
//                    |
//     Category       | A categorization of the content of this package.
//                    |
//     Version        | The version number. This value is set by the user or by the application.
//
// For example:
//
//    err := f.SetDocProps(&excelize.DocProperties{
//        Category:       "category",
//        ContentStatus:  "Draft",
//        Created:        "2019-06-04T22:00:10Z",
//        Creator:        "Go Excelize",
//        Description:    "This file created by Go Excelize",
//        Identifier:     "xlsx",
//        Keywords:       "Spreadsheet",
//        LastModifiedBy: "Go Author",
//        Modified:       "2019-06-04T22:00:10Z",
//        Revision:       "0",
//        Subject:        "Test Subject",
//        Title:          "Test Title",
//        Language:       "en-US",
//        Version:        "1.0.0",
//    })
//
func (f *File) SetDocProps(docProperties *DocProperties) (err error) {
	var (
		core               *decodeCoreProperties
		newProps           *xlsxCoreProperties
		fields             []string
		output             []byte
		immutable, mutable reflect.Value
		field, val         string
	)

	core = new(decodeCoreProperties)
	if err = f.xmlNewDecoder(bytes.NewReader(namespaceStrictToTransitional(f.readXML(defaultXMLPathDocPropsCore)))).
		Decode(core); err != nil && err != io.EOF {
		err = fmt.Errorf("xml decode error: %s", err)
		return
	}
	newProps, err = &xlsxCoreProperties{
		Dc:             NameSpaceDublinCore,
		Dcterms:        NameSpaceDublinCoreTerms,
		Dcmitype:       NameSpaceDublinCoreMetadataInitiative,
		XSI:            NameSpaceXMLSchemaInstance,
		Title:          core.Title,
		Subject:        core.Subject,
		Creator:        core.Creator,
		Keywords:       core.Keywords,
		Description:    core.Description,
		LastModifiedBy: core.LastModifiedBy,
		Language:       core.Language,
		Identifier:     core.Identifier,
		Revision:       core.Revision,
		ContentStatus:  core.ContentStatus,
		Category:       core.Category,
		Version:        core.Version,
	}, nil
	newProps.Created.Text, newProps.Created.Type, newProps.Modified.Text, newProps.Modified.Type = core.Created.Text, core.Created.Type, core.Modified.Text, core.Modified.Type
	fields = []string{
		"Category", "ContentStatus", "Creator", "Description", "Identifier", "Keywords",
		"LastModifiedBy", "Revision", "Subject", "Title", "Language", "Version",
	}
	immutable, mutable = reflect.ValueOf(*docProperties), reflect.ValueOf(newProps).Elem()
	for _, field = range fields {
		if val = immutable.FieldByName(field).String(); val != "" {
			mutable.FieldByName(field).SetString(val)
		}
	}
	if docProperties.Created != "" {
		newProps.Created.Text = docProperties.Created
	}
	if docProperties.Modified != "" {
		newProps.Modified.Text = docProperties.Modified
	}
	output, err = xml.Marshal(newProps)
	f.saveFileList(defaultXMLPathDocPropsCore, output)

	return
}

// GetDocProps provides a function to get document core properties.
func (f *File) GetDocProps() (ret *DocProperties, err error) {
	core := new(decodeCoreProperties)

	if err = f.xmlNewDecoder(bytes.NewReader(namespaceStrictToTransitional(f.readXML(defaultXMLPathDocPropsCore)))).
		Decode(core); err != nil && err != io.EOF {
		err = fmt.Errorf("xml decode error: %s", err)
		return
	}
	ret, err = &DocProperties{
		Category:       core.Category,
		ContentStatus:  core.ContentStatus,
		Created:        core.Created.Text,
		Creator:        core.Creator,
		Description:    core.Description,
		Identifier:     core.Identifier,
		Keywords:       core.Keywords,
		LastModifiedBy: core.LastModifiedBy,
		Modified:       core.Modified.Text,
		Revision:       core.Revision,
		Subject:        core.Subject,
		Title:          core.Title,
		Language:       core.Language,
		Version:        core.Version,
	}, nil

	return
}
