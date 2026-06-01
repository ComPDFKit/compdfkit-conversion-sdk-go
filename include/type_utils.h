#ifndef TYPE_UTILS_H
#define TYPE_UTILS_H

#include <algorithm>
#include <cstring>

#include "compdf_basictypes_c.h"
#include "common/base_type.h"

namespace compdf::utils
{

#define CERRORCODE(x) static_cast<CSDKErrorCode>(x)

inline
base::ConvertOptions ConvertOptionType(CConvertOption option)
{
    base::ConvertOptions ret;

    ret.enable_ai_layout = option.enable_ai_layout;
    ret.enable_ai_table_recognition = option.enable_ai_table_recognition;
    ret.contain_image = option.contain_image;
    ret.contain_page_background_image = option.contain_page_background_image;
    ret.json_contain_table = option.json_contain_table;
    ret.contain_annotation = option.contain_annotation;
    ret.excel_all_content = option.excel_all_content;
    ret.excel_csv_format = option.excel_csv_format;
    ret.enable_ocr = option.enable_ocr;
    ret.transparent_text = option.transparent_text;
    ret.txt_table_format = option.txt_table_format;
    ret.image_path_enhance = option.image_path_enhance;
    ret.formula_to_image = option.formula_to_image;
    ret.auto_create_folder = option.auto_create_folder;
    ret.output_document_per_page = option.output_document_per_page;
    ret.image_scaling = option.image_scaling;
    ret.page_layout_mode = static_cast<base::PageLayoutMode>(option.page_layout_mode);
    ret.excel_worksheet_option = static_cast<base::ExcelWorksheetOption>(option.excel_worksheet_option);
    ret.html_option = static_cast<base::HtmlOption>(option.html_option);
    ret.image_color_mode = static_cast<base::ImageColorMode>(option.image_color_mode);
    ret.image_type = static_cast<base::ImageType>(option.image_type);
    ret.ocr_option = static_cast<base::OCROption>(option.ocr_option);
    std::memcpy(ret.font_name, option.font_name, sizeof(ret.font_name));
    ret.font_name[sizeof(ret.font_name) - 1] = '\0';
    std::memcpy(ret.page_ranges, option.page_ranges, sizeof(ret.page_ranges));
    ret.page_ranges[sizeof(ret.page_ranges) - 1] = '\0';

    ret.language_count = std::min(option.language_count, 32);
    for (int i = 0; i < ret.language_count; i++)
        ret.languages[i] = static_cast<base::OCRLanguage>(option.languages[i]);

    return ret;
}

}


#endif //TYPE_UTILS_H
