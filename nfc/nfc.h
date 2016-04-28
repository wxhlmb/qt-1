#pragma once

#ifndef GO_QTNFC_H
#define GO_QTNFC_H

#ifdef __cplusplus
extern "C" {
#endif

void* QNdefFilter_NewQNdefFilter();
void* QNdefFilter_NewQNdefFilter2(void* other);
void QNdefFilter_Clear(void* ptr);
int QNdefFilter_OrderMatch(void* ptr);
int QNdefFilter_RecordCount(void* ptr);
void QNdefFilter_SetOrderMatch(void* ptr, int on);
void QNdefFilter_DestroyQNdefFilter(void* ptr);
void* QNdefMessage_NewQNdefMessage();
void* QNdefMessage_NewQNdefMessage3(void* message);
void* QNdefMessage_NewQNdefMessage2(void* record);
void* QNdefMessage_QNdefMessage_FromByteArray(char* message);
char* QNdefMessage_ToByteArray(void* ptr);
void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord();
void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord3(void* other);
void* QNdefNfcSmartPosterRecord_NewQNdefNfcSmartPosterRecord2(void* other);
int QNdefNfcSmartPosterRecord_Action(void* ptr);
void QNdefNfcSmartPosterRecord_AddIcon2(void* ptr, char* ty, char* data);
int QNdefNfcSmartPosterRecord_AddTitle(void* ptr, void* text);
int QNdefNfcSmartPosterRecord_AddTitle2(void* ptr, char* text, char* locale, int encoding);
int QNdefNfcSmartPosterRecord_HasAction(void* ptr);
int QNdefNfcSmartPosterRecord_HasIcon(void* ptr, char* mimetype);
int QNdefNfcSmartPosterRecord_HasSize(void* ptr);
int QNdefNfcSmartPosterRecord_HasTitle(void* ptr, char* locale);
int QNdefNfcSmartPosterRecord_HasTypeInfo(void* ptr);
char* QNdefNfcSmartPosterRecord_Icon(void* ptr, char* mimetype);
int QNdefNfcSmartPosterRecord_IconCount(void* ptr);
int QNdefNfcSmartPosterRecord_RemoveIcon2(void* ptr, char* ty);
int QNdefNfcSmartPosterRecord_RemoveTitle(void* ptr, void* text);
int QNdefNfcSmartPosterRecord_RemoveTitle2(void* ptr, char* locale);
void QNdefNfcSmartPosterRecord_SetAction(void* ptr, int act);
void QNdefNfcSmartPosterRecord_SetTypeInfo(void* ptr, char* ty);
void QNdefNfcSmartPosterRecord_SetUri(void* ptr, void* url);
void QNdefNfcSmartPosterRecord_SetUri2(void* ptr, void* url);
char* QNdefNfcSmartPosterRecord_Title(void* ptr, char* locale);
int QNdefNfcSmartPosterRecord_TitleCount(void* ptr);
char* QNdefNfcSmartPosterRecord_TypeInfo(void* ptr);
void* QNdefNfcSmartPosterRecord_Uri(void* ptr);
void QNdefNfcSmartPosterRecord_DestroyQNdefNfcSmartPosterRecord(void* ptr);
void* QNdefNfcTextRecord_NewQNdefNfcTextRecord();
void* QNdefNfcTextRecord_NewQNdefNfcTextRecord2(void* other);
int QNdefNfcTextRecord_Encoding(void* ptr);
char* QNdefNfcTextRecord_Locale(void* ptr);
void QNdefNfcTextRecord_SetEncoding(void* ptr, int encoding);
void QNdefNfcTextRecord_SetLocale(void* ptr, char* locale);
void QNdefNfcTextRecord_SetText(void* ptr, char* text);
char* QNdefNfcTextRecord_Text(void* ptr);
void* QNdefNfcUriRecord_NewQNdefNfcUriRecord();
void* QNdefNfcUriRecord_NewQNdefNfcUriRecord2(void* other);
void QNdefNfcUriRecord_SetUri(void* ptr, void* uri);
void* QNdefNfcUriRecord_Uri(void* ptr);
void* QNdefRecord_NewQNdefRecord();
void* QNdefRecord_NewQNdefRecord2(void* other);
char* QNdefRecord_Id(void* ptr);
int QNdefRecord_IsEmpty(void* ptr);
char* QNdefRecord_Payload(void* ptr);
void QNdefRecord_SetId(void* ptr, char* id);
void QNdefRecord_SetPayload(void* ptr, char* payload);
void QNdefRecord_SetType(void* ptr, char* ty);
void QNdefRecord_SetTypeNameFormat(void* ptr, int typeNameFormat);
char* QNdefRecord_Type(void* ptr);
int QNdefRecord_TypeNameFormat(void* ptr);
void QNdefRecord_DestroyQNdefRecord(void* ptr);
int QNearFieldManager_RegisterNdefMessageHandler(void* ptr, void* object, char* method);
int QNearFieldManager_StartTargetDetection(void* ptr);
void* QNearFieldManager_NewQNearFieldManager(void* parent);
int QNearFieldManager_IsAvailable(void* ptr);
int QNearFieldManager_RegisterNdefMessageHandler2(void* ptr, int typeNameFormat, char* ty, void* object, char* method);
int QNearFieldManager_RegisterNdefMessageHandler3(void* ptr, void* filter, void* object, char* method);
void QNearFieldManager_SetTargetAccessModes(void* ptr, int accessModes);
void QNearFieldManager_StopTargetDetection(void* ptr);
int QNearFieldManager_TargetAccessModes(void* ptr);
void QNearFieldManager_ConnectTargetDetected(void* ptr);
void QNearFieldManager_DisconnectTargetDetected(void* ptr);
void QNearFieldManager_TargetDetected(void* ptr, void* target);
void QNearFieldManager_ConnectTargetLost(void* ptr);
void QNearFieldManager_DisconnectTargetLost(void* ptr);
void QNearFieldManager_TargetLost(void* ptr, void* target);
int QNearFieldManager_UnregisterNdefMessageHandler(void* ptr, int handlerId);
void QNearFieldManager_DestroyQNearFieldManager(void* ptr);
void QNearFieldManager_TimerEvent(void* ptr, void* event);
void QNearFieldManager_TimerEventDefault(void* ptr, void* event);
void QNearFieldManager_ChildEvent(void* ptr, void* event);
void QNearFieldManager_ChildEventDefault(void* ptr, void* event);
void QNearFieldManager_ConnectNotify(void* ptr, void* sign);
void QNearFieldManager_ConnectNotifyDefault(void* ptr, void* sign);
void QNearFieldManager_CustomEvent(void* ptr, void* event);
void QNearFieldManager_CustomEventDefault(void* ptr, void* event);
void QNearFieldManager_DeleteLater(void* ptr);
void QNearFieldManager_DisconnectNotify(void* ptr, void* sign);
void QNearFieldManager_DisconnectNotifyDefault(void* ptr, void* sign);
int QNearFieldManager_Event(void* ptr, void* e);
int QNearFieldManager_EventDefault(void* ptr, void* e);
int QNearFieldManager_EventFilter(void* ptr, void* watched, void* event);
int QNearFieldManager_EventFilterDefault(void* ptr, void* watched, void* event);
void* QNearFieldManager_MetaObject(void* ptr);
void* QNearFieldManager_MetaObjectDefault(void* ptr);
void* QNearFieldShareManager_NewQNearFieldShareManager(void* parent);
void QNearFieldShareManager_ConnectError(void* ptr);
void QNearFieldShareManager_DisconnectError(void* ptr);
void QNearFieldShareManager_Error(void* ptr, int error);
void QNearFieldShareManager_SetShareModes(void* ptr, int mode);
int QNearFieldShareManager_ShareError(void* ptr);
int QNearFieldShareManager_ShareModes(void* ptr);
void QNearFieldShareManager_ConnectShareModesChanged(void* ptr);
void QNearFieldShareManager_DisconnectShareModesChanged(void* ptr);
void QNearFieldShareManager_ShareModesChanged(void* ptr, int modes);
int QNearFieldShareManager_QNearFieldShareManager_SupportedShareModes();
void QNearFieldShareManager_ConnectTargetDetected(void* ptr);
void QNearFieldShareManager_DisconnectTargetDetected(void* ptr);
void QNearFieldShareManager_TargetDetected(void* ptr, void* shareTarget);
void QNearFieldShareManager_DestroyQNearFieldShareManager(void* ptr);
void QNearFieldShareManager_TimerEvent(void* ptr, void* event);
void QNearFieldShareManager_TimerEventDefault(void* ptr, void* event);
void QNearFieldShareManager_ChildEvent(void* ptr, void* event);
void QNearFieldShareManager_ChildEventDefault(void* ptr, void* event);
void QNearFieldShareManager_ConnectNotify(void* ptr, void* sign);
void QNearFieldShareManager_ConnectNotifyDefault(void* ptr, void* sign);
void QNearFieldShareManager_CustomEvent(void* ptr, void* event);
void QNearFieldShareManager_CustomEventDefault(void* ptr, void* event);
void QNearFieldShareManager_DeleteLater(void* ptr);
void QNearFieldShareManager_DisconnectNotify(void* ptr, void* sign);
void QNearFieldShareManager_DisconnectNotifyDefault(void* ptr, void* sign);
int QNearFieldShareManager_Event(void* ptr, void* e);
int QNearFieldShareManager_EventDefault(void* ptr, void* e);
int QNearFieldShareManager_EventFilter(void* ptr, void* watched, void* event);
int QNearFieldShareManager_EventFilterDefault(void* ptr, void* watched, void* event);
void* QNearFieldShareManager_MetaObject(void* ptr);
void* QNearFieldShareManager_MetaObjectDefault(void* ptr);
void QNearFieldShareTarget_Cancel(void* ptr);
void QNearFieldShareTarget_ConnectError(void* ptr);
void QNearFieldShareTarget_DisconnectError(void* ptr);
void QNearFieldShareTarget_Error(void* ptr, int error);
int QNearFieldShareTarget_IsShareInProgress(void* ptr);
int QNearFieldShareTarget_Share(void* ptr, void* message);
int QNearFieldShareTarget_ShareError(void* ptr);
void QNearFieldShareTarget_ConnectShareFinished(void* ptr);
void QNearFieldShareTarget_DisconnectShareFinished(void* ptr);
void QNearFieldShareTarget_ShareFinished(void* ptr);
int QNearFieldShareTarget_ShareModes(void* ptr);
void QNearFieldShareTarget_DestroyQNearFieldShareTarget(void* ptr);
void QNearFieldShareTarget_TimerEvent(void* ptr, void* event);
void QNearFieldShareTarget_TimerEventDefault(void* ptr, void* event);
void QNearFieldShareTarget_ChildEvent(void* ptr, void* event);
void QNearFieldShareTarget_ChildEventDefault(void* ptr, void* event);
void QNearFieldShareTarget_ConnectNotify(void* ptr, void* sign);
void QNearFieldShareTarget_ConnectNotifyDefault(void* ptr, void* sign);
void QNearFieldShareTarget_CustomEvent(void* ptr, void* event);
void QNearFieldShareTarget_CustomEventDefault(void* ptr, void* event);
void QNearFieldShareTarget_DeleteLater(void* ptr);
void QNearFieldShareTarget_DisconnectNotify(void* ptr, void* sign);
void QNearFieldShareTarget_DisconnectNotifyDefault(void* ptr, void* sign);
int QNearFieldShareTarget_Event(void* ptr, void* e);
int QNearFieldShareTarget_EventDefault(void* ptr, void* e);
int QNearFieldShareTarget_EventFilter(void* ptr, void* watched, void* event);
int QNearFieldShareTarget_EventFilterDefault(void* ptr, void* watched, void* event);
void* QNearFieldShareTarget_MetaObject(void* ptr);
void* QNearFieldShareTarget_MetaObjectDefault(void* ptr);
void* QNearFieldTarget_NewQNearFieldTarget(void* parent);
int QNearFieldTarget_AccessMethods(void* ptr);
void QNearFieldTarget_ConnectDisconnected(void* ptr);
void QNearFieldTarget_DisconnectDisconnected(void* ptr);
void QNearFieldTarget_Disconnected(void* ptr);
int QNearFieldTarget_HasNdefMessage(void* ptr);
int QNearFieldTarget_HasNdefMessageDefault(void* ptr);
int QNearFieldTarget_IsProcessingCommand(void* ptr);
void QNearFieldTarget_ConnectNdefMessageRead(void* ptr);
void QNearFieldTarget_DisconnectNdefMessageRead(void* ptr);
void QNearFieldTarget_NdefMessageRead(void* ptr, void* message);
void QNearFieldTarget_ConnectNdefMessagesWritten(void* ptr);
void QNearFieldTarget_DisconnectNdefMessagesWritten(void* ptr);
void QNearFieldTarget_NdefMessagesWritten(void* ptr);
int QNearFieldTarget_Type(void* ptr);
char* QNearFieldTarget_Uid(void* ptr);
void* QNearFieldTarget_Url(void* ptr);
void* QNearFieldTarget_UrlDefault(void* ptr);
void QNearFieldTarget_DestroyQNearFieldTarget(void* ptr);
void QNearFieldTarget_TimerEvent(void* ptr, void* event);
void QNearFieldTarget_TimerEventDefault(void* ptr, void* event);
void QNearFieldTarget_ChildEvent(void* ptr, void* event);
void QNearFieldTarget_ChildEventDefault(void* ptr, void* event);
void QNearFieldTarget_ConnectNotify(void* ptr, void* sign);
void QNearFieldTarget_ConnectNotifyDefault(void* ptr, void* sign);
void QNearFieldTarget_CustomEvent(void* ptr, void* event);
void QNearFieldTarget_CustomEventDefault(void* ptr, void* event);
void QNearFieldTarget_DeleteLater(void* ptr);
void QNearFieldTarget_DisconnectNotify(void* ptr, void* sign);
void QNearFieldTarget_DisconnectNotifyDefault(void* ptr, void* sign);
int QNearFieldTarget_Event(void* ptr, void* e);
int QNearFieldTarget_EventDefault(void* ptr, void* e);
int QNearFieldTarget_EventFilter(void* ptr, void* watched, void* event);
int QNearFieldTarget_EventFilterDefault(void* ptr, void* watched, void* event);
void* QNearFieldTarget_MetaObject(void* ptr);
void* QNearFieldTarget_MetaObjectDefault(void* ptr);
int QQmlNdefRecord_TypeNameFormat(void* ptr);
void* QQmlNdefRecord_NewQQmlNdefRecord(void* parent);
void* QQmlNdefRecord_NewQQmlNdefRecord2(void* record, void* parent);
void* QQmlNdefRecord_Record(void* ptr);
void QQmlNdefRecord_ConnectRecordChanged(void* ptr);
void QQmlNdefRecord_DisconnectRecordChanged(void* ptr);
void QQmlNdefRecord_RecordChanged(void* ptr);
void QQmlNdefRecord_SetRecord(void* ptr, void* record);
void QQmlNdefRecord_SetType(void* ptr, char* newtype);
void QQmlNdefRecord_SetTypeNameFormat(void* ptr, int newTypeNameFormat);
char* QQmlNdefRecord_Type(void* ptr);
void QQmlNdefRecord_ConnectTypeChanged(void* ptr);
void QQmlNdefRecord_DisconnectTypeChanged(void* ptr);
void QQmlNdefRecord_TypeChanged(void* ptr);
void QQmlNdefRecord_ConnectTypeNameFormatChanged(void* ptr);
void QQmlNdefRecord_DisconnectTypeNameFormatChanged(void* ptr);
void QQmlNdefRecord_TypeNameFormatChanged(void* ptr);
void QQmlNdefRecord_TimerEvent(void* ptr, void* event);
void QQmlNdefRecord_TimerEventDefault(void* ptr, void* event);
void QQmlNdefRecord_ChildEvent(void* ptr, void* event);
void QQmlNdefRecord_ChildEventDefault(void* ptr, void* event);
void QQmlNdefRecord_ConnectNotify(void* ptr, void* sign);
void QQmlNdefRecord_ConnectNotifyDefault(void* ptr, void* sign);
void QQmlNdefRecord_CustomEvent(void* ptr, void* event);
void QQmlNdefRecord_CustomEventDefault(void* ptr, void* event);
void QQmlNdefRecord_DeleteLater(void* ptr);
void QQmlNdefRecord_DisconnectNotify(void* ptr, void* sign);
void QQmlNdefRecord_DisconnectNotifyDefault(void* ptr, void* sign);
int QQmlNdefRecord_Event(void* ptr, void* e);
int QQmlNdefRecord_EventDefault(void* ptr, void* e);
int QQmlNdefRecord_EventFilter(void* ptr, void* watched, void* event);
int QQmlNdefRecord_EventFilterDefault(void* ptr, void* watched, void* event);
void* QQmlNdefRecord_MetaObject(void* ptr);
void* QQmlNdefRecord_MetaObjectDefault(void* ptr);

#ifdef __cplusplus
}
#endif

#endif