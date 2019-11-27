/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_Channel = map[string]string{
	"":     "Channel represents a message transmission channel in TKE.",
	"spec": "Spec defines the desired channel.",
}

func (Channel) SwaggerDoc() map[string]string {
	return map_Channel
}

var map_ChannelList = map[string]string{
	"":      "ChannelList is the whole list of all channels which owned by a tenant.",
	"items": "List of channels.",
}

func (ChannelList) SwaggerDoc() map[string]string {
	return map_ChannelList
}

var map_ChannelSMTP = map[string]string{
	"": "ChannelSMTP indicates a channel configuration for sending email notifications using the SMTP server.",
}

func (ChannelSMTP) SwaggerDoc() map[string]string {
	return map_ChannelSMTP
}

var map_ChannelSpec = map[string]string{
	"":           "ChannelSpec is a description of a channel.",
	"finalizers": "Finalizers is an opaque list of values that must be empty to permanently remove object from storage.",
}

func (ChannelSpec) SwaggerDoc() map[string]string {
	return map_ChannelSpec
}

var map_ChannelStatus = map[string]string{
	"": "ChannelStatus represents information about the status of a cluster.",
}

func (ChannelStatus) SwaggerDoc() map[string]string {
	return map_ChannelStatus
}

var map_ChannelTencentCloudSMS = map[string]string{
	"": "ChannelTencentCloudSMS indicates the channel configuration for sending messages using Tencent Cloud SMS Gateway. See: https://cloud.tencent.com/document/product/382/5976",
}

func (ChannelTencentCloudSMS) SwaggerDoc() map[string]string {
	return map_ChannelTencentCloudSMS
}

var map_ChannelWechat = map[string]string{
	"":      "ChannelWechat indicates a channel configuration for sending template notifications using WeChat.",
	"appID": "AppID indicates the unique credentials of the third-party user. See https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183",
}

func (ChannelWechat) SwaggerDoc() map[string]string {
	return map_ChannelWechat
}

var map_ConfigMap = map[string]string{
	"":           "ConfigMap holds configuration data for tke to consume.",
	"data":       "Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.",
	"binaryData": "BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process.",
}

func (ConfigMap) SwaggerDoc() map[string]string {
	return map_ConfigMap
}

var map_ConfigMapList = map[string]string{
	"":      "ConfigMapList is a resource containing a list of ConfigMap objects.",
	"items": "Items is the list of ConfigMaps.",
}

func (ConfigMapList) SwaggerDoc() map[string]string {
	return map_ConfigMapList
}

var map_Message = map[string]string{
	"":     "Message indicates a message in the notification system.",
	"spec": "Spec defines the desired message.",
}

func (Message) SwaggerDoc() map[string]string {
	return map_Message
}

var map_MessageList = map[string]string{
	"":      "MessageList is the whole list of all message which owned by a tenant.",
	"items": "List of messages.",
}

func (MessageList) SwaggerDoc() map[string]string {
	return map_MessageList
}

var map_MessageRequest = map[string]string{
	"":     "MessageRequest represents a message sending request, which may include multiple recipients and multiple receiving groups.",
	"spec": "Spec defines the desired message.",
}

func (MessageRequest) SwaggerDoc() map[string]string {
	return map_MessageRequest
}

var map_MessageRequestList = map[string]string{
	"":      "MessageRequestList is the whole list of all message which owned by a tenant.",
	"items": "List of message requests.",
}

func (MessageRequestList) SwaggerDoc() map[string]string {
	return map_MessageRequestList
}

var map_MessageRequestSpec = map[string]string{
	"": "MessageRequestSpec is a description of a message request.",
}

func (MessageRequestSpec) SwaggerDoc() map[string]string {
	return map_MessageRequestSpec
}

var map_MessageRequestStatus = map[string]string{
	"":                   "MessageRequestStatus represents information about the status of a message request.",
	"lastTransitionTime": "The last time the condition transitioned from one status to another.",
	"errors":             "A human readable message indicating details about the transition.",
}

func (MessageRequestStatus) SwaggerDoc() map[string]string {
	return map_MessageRequestStatus
}

var map_MessageSpec = map[string]string{
	"": "MessageSpec is a description of a message.",
}

func (MessageSpec) SwaggerDoc() map[string]string {
	return map_MessageSpec
}

var map_MessageStatus = map[string]string{
	"":                   "MessageStatus represents information about the status of a message.",
	"lastTransitionTime": "The last time the condition transitioned from one status to another.",
}

func (MessageStatus) SwaggerDoc() map[string]string {
	return map_MessageStatus
}

var map_Receiver = map[string]string{
	"":     "Receiver indicates a message notification recipient, usually representing a user in the user system or a webhook service address.",
	"spec": "Spec defines the desired receiver.",
}

func (Receiver) SwaggerDoc() map[string]string {
	return map_Receiver
}

var map_ReceiverGroup = map[string]string{
	"":     "ReceiverGroup indicates multiple message recipients.",
	"spec": "Spec defines the desired receiver group.",
}

func (ReceiverGroup) SwaggerDoc() map[string]string {
	return map_ReceiverGroup
}

var map_ReceiverGroupList = map[string]string{
	"":      "ReceiverGroupList is the whole list of all receiver which owned by a tenant.",
	"items": "List of receiver groups.",
}

func (ReceiverGroupList) SwaggerDoc() map[string]string {
	return map_ReceiverGroupList
}

var map_ReceiverGroupSpec = map[string]string{
	"": "ReceiverGroupSpec is a description of a receiver group.",
}

func (ReceiverGroupSpec) SwaggerDoc() map[string]string {
	return map_ReceiverGroupSpec
}

var map_ReceiverList = map[string]string{
	"":      "ReceiverList is the whole list of all receiver which owned by a tenant.",
	"items": "List of receivers.",
}

func (ReceiverList) SwaggerDoc() map[string]string {
	return map_ReceiverList
}

var map_ReceiverSpec = map[string]string{
	"":           "ReceiverSpec is a description of a receiver.",
	"identities": "Identities represents the characteristics of the message recipient. The hash table key represents the message delivery channel id, and the value represents the user identification number in the channel. For example, if it is a short message sending channel, then the value is the user's mobile phone number; if it is a mail sending channel, then the value is the user's email address.",
}

func (ReceiverSpec) SwaggerDoc() map[string]string {
	return map_ReceiverSpec
}

var map_Template = map[string]string{
	"":     "Template indicates the template used to send notifications under this channel.",
	"spec": "Spec defines the desired template.",
}

func (Template) SwaggerDoc() map[string]string {
	return map_Template
}

var map_TemplateList = map[string]string{
	"":      "TemplateList is the whole list of all template which owned by a channel.",
	"items": "List of templates.",
}

func (TemplateList) SwaggerDoc() map[string]string {
	return map_TemplateList
}

var map_TemplateSpec = map[string]string{
	"": "TemplateSpec is a description of a template.",
}

func (TemplateSpec) SwaggerDoc() map[string]string {
	return map_TemplateSpec
}

var map_TemplateTencentCloudSMS = map[string]string{
	"": "TemplateTencentCloudSMS indicates the template used when sending text messages using Tencent Cloud SMS Gateway. The template must be approved.",
}

func (TemplateTencentCloudSMS) SwaggerDoc() map[string]string {
	return map_TemplateTencentCloudSMS
}

var map_TemplateText = map[string]string{
	"": "TemplateText indicates the template used to send notifications using other non-templated channels.",
}

func (TemplateText) SwaggerDoc() map[string]string {
	return map_TemplateText
}

var map_TemplateWechat = map[string]string{
	"":                 "TemplateWechat indicates the template when sending a text message using the WeChat public account. The template must be approved and registered.",
	"templateID":       "TemplateID indicates the template id of the template message notification. See https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140183",
	"url":              "URL indicates the web address of the user who clicked the notification in WeChat.",
	"miniProgramAppID": "MiniProgramAppID indicates the unique identification number of the WeChat applet that the user clicked on the notification in WeChat.",
}

func (TemplateWechat) SwaggerDoc() map[string]string {
	return map_TemplateWechat
}

// AUTO-GENERATED FUNCTIONS END HERE