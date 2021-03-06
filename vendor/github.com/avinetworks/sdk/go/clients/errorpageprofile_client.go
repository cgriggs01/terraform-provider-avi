/***************************************************************************
 *
 * AVI CONFIDENTIAL
 * __________________
 *
 * [2013] - [2018] Avi Networks Incorporated
 * All Rights Reserved.
 *
 * NOTICE: All information contained herein is, and remains the property
 * of Avi Networks Incorporated and its suppliers, if any. The intellectual
 * and technical concepts contained herein are proprietary to Avi Networks
 * Incorporated, and its suppliers and are covered by U.S. and Foreign
 * Patents, patents in process, and are protected by trade secret or
 * copyright law, and other laws. Dissemination of this information or
 * reproduction of this material is strictly forbidden unless prior written
 * permission is obtained from Avi Networks Incorporated.
 */

package clients

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

import (
	"github.com/avinetworks/sdk/go/models"
	"github.com/avinetworks/sdk/go/session"
)

// ErrorPageProfileClient is a client for avi ErrorPageProfile resource
type ErrorPageProfileClient struct {
	aviSession *session.AviSession
}

// NewErrorPageProfileClient creates a new client for ErrorPageProfile resource
func NewErrorPageProfileClient(aviSession *session.AviSession) *ErrorPageProfileClient {
	return &ErrorPageProfileClient{aviSession: aviSession}
}

func (client *ErrorPageProfileClient) getAPIPath(uuid string) string {
	path := "api/errorpageprofile"
	if uuid != "" {
		path += "/" + uuid
	}
	return path
}

// GetAll is a collection API to get a list of ErrorPageProfile objects
func (client *ErrorPageProfileClient) GetAll() ([]*models.ErrorPageProfile, error) {
	var plist []*models.ErrorPageProfile
	err := client.aviSession.GetCollection(client.getAPIPath(""), &plist)
	return plist, err
}

// Get an existing ErrorPageProfile by uuid
func (client *ErrorPageProfileClient) Get(uuid string) (*models.ErrorPageProfile, error) {
	var obj *models.ErrorPageProfile
	err := client.aviSession.Get(client.getAPIPath(uuid), &obj)
	return obj, err
}

// GetByName - Get an existing ErrorPageProfile by name
func (client *ErrorPageProfileClient) GetByName(name string) (*models.ErrorPageProfile, error) {
	var obj *models.ErrorPageProfile
	err := client.aviSession.GetObjectByName("errorpageprofile", name, &obj)
	return obj, err
}

// GetObject - Get an existing ErrorPageProfile by filters like name, cloud, tenant
// Api creates ErrorPageProfile object with every call.
func (client *ErrorPageProfileClient) GetObject(options ...session.ApiOptionsParams) (*models.ErrorPageProfile, error) {
	var obj *models.ErrorPageProfile
	newOptions := make([]session.ApiOptionsParams, len(options)+1)
	for i, p := range options {
		newOptions[i] = p
	}
	newOptions[len(options)] = session.SetResult(&obj)
	err := client.aviSession.GetObject("errorpageprofile", newOptions...)
	return obj, err
}

// Create a new ErrorPageProfile object
func (client *ErrorPageProfileClient) Create(obj *models.ErrorPageProfile) (*models.ErrorPageProfile, error) {
	var robj *models.ErrorPageProfile
	err := client.aviSession.Post(client.getAPIPath(""), obj, &robj)
	return robj, err
}

// Update an existing ErrorPageProfile object
func (client *ErrorPageProfileClient) Update(obj *models.ErrorPageProfile) (*models.ErrorPageProfile, error) {
	var robj *models.ErrorPageProfile
	path := client.getAPIPath(*obj.UUID)
	err := client.aviSession.Put(path, obj, &robj)
	return robj, err
}

// Delete an existing ErrorPageProfile object with a given UUID
func (client *ErrorPageProfileClient) Delete(uuid string) error {
	return client.aviSession.Delete(client.getAPIPath(uuid))
}

// DeleteByName - Delete an existing ErrorPageProfile object with a given name
func (client *ErrorPageProfileClient) DeleteByName(name string) error {
	res, err := client.GetByName(name)
	if err != nil {
		return err
	}
	return client.Delete(*res.UUID)
}

// GetAviSession
func (client *ErrorPageProfileClient) GetAviSession() *session.AviSession {
	return client.aviSession
}
