#ifndef COMPDFKIT_DOCUMENT_AI_C_H
#define COMPDFKIT_DOCUMENT_AI_C_H

#ifdef __cplusplus
extern "C" {
#endif

#include "compdf_basictypes_c.h"
#include "sdk_document_ai_common.h"

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_Ocr(COMPDFKIT_STRING file_path,
										   COCRLanguage* languages,
										   int languagecount,
										   da_ocr_det_t** det_result,
										   da_ocr_rec_t** rec_result,
										   int* result_count);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_Ocr0(void* image_data,
											int data_length,
											COCRLanguage* languages,
											int languagecount,
											da_ocr_det_t** det_result,
											da_ocr_rec_t** rec_result,
											int* result_count);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_LayoutAnalysis(COMPDFKIT_STRING file_path, da_detection_t** detection_result, int* detection_count);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_LayoutAnalysis0(void* image_data, int data_length, da_detection_t** detection_result, int* detection_count);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_StampDetection(COMPDFKIT_STRING file_path, da_detection_t** detection_result, int* detection_count);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_StampDetection0(void* image_data, int data_length, da_detection_t** detection_result, int* detection_count);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_TableRec(COMPDFKIT_STRING file_path, da_table_t** table_result, int* table_count);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_TableRec0(void* image_data, int data_length, da_table_t** table_result, int* table_count);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_MagicColor(COMPDFKIT_STRING file_path, COMPDFKIT_STRING output_path);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_MagicColor0(void* image_data, int data_length, COMPDFKIT_STRING output_path);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_MagicColorToImage(COMPDFKIT_STRING file_path, da_image_t* output_img);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_MagicColorToImage0(void* image_data, int data_length, da_image_t* output_img);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_Dewarp(COMPDFKIT_STRING file_path, da_dewarp_t* doc, COMPDFKIT_STRING output_path);

COMPDFKIT_DECL
CSDKErrorCode CSDK_COMPDFKIT_CALL CPDF_Dewarp0(void* image_data, int data_length, da_dewarp_t* doc, COMPDFKIT_STRING output_path);

COMPDFKIT_DECL
void CSDK_COMPDFKIT_CALL CPDF_OcrRelease(da_ocr_det_t* det_result,da_ocr_rec_t* rec_result);

COMPDFKIT_DECL
void CSDK_COMPDFKIT_CALL CPDF_LayoutAnalysisRelease(da_detection_t* detection_result);

COMPDFKIT_DECL
void CSDK_COMPDFKIT_CALL CPDF_TableRecRelease(da_table_t** table_result, int table_count);

COMPDFKIT_DECL
void CSDK_COMPDFKIT_CALL CPDF_MagicColorRelease(da_image_t* image);

COMPDFKIT_DECL
void CSDK_COMPDFKIT_CALL CPDF_DewarpRelease(da_dewarp_t* doc);

COMPDFKIT_DECL
void CSDK_COMPDFKIT_CALL CPDF_StampDetectionRelease(da_detection_t* detection_result);


#ifdef __cplusplus
}
#endif
#endif //COMPDFKIT_DOCUMENT_AI_C_H
