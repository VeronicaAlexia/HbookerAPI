import pydantic
import typing


class LastChapterInfo(pydantic.BaseModel):
    chapter_id: str
    book_id: str
    chapter_index: typing.Optional[str] = None
    chapter_title: typing.Optional[str] = None
    uptime: typing.Optional[str] = None
    mtime: typing.Optional[str] = None


class TagList(pydantic.BaseModel):
    tag_id: typing.Optional[str] = None
    tag_type: typing.Optional[str] = None
    tag_name: typing.Optional[str] = None


class BookInfo(pydantic.BaseModel):
    book_id: typing.Optional[str] = None
    book_name: typing.Optional[str] = None
    description: typing.Optional[str] = None
    book_src: typing.Optional[str] = None
    category_index: typing.Optional[str] = None
    tag: typing.Optional[str] = None
    total_word_count: typing.Optional[str] = None
    up_status: typing.Optional[str] = None
    update_status: typing.Optional[str] = None
    is_paid: typing.Optional[str] = None
    discount: typing.Optional[str] = None
    discount_end_time: typing.Optional[str] = None
    cover: typing.Optional[str] = None
    author_name: typing.Optional[str] = None
    uptime: typing.Optional[str] = None
    newtime: typing.Optional[str] = None
    review_amount: typing.Optional[str] = None
    reward_amount: typing.Optional[str] = None
    chapter_amount: typing.Optional[str] = None
    is_original: typing.Optional[str] = None
    total_click: typing.Optional[str] = None
    month_click: typing.Optional[str] = None
    week_click: typing.Optional[str] = None
    month_no_vip_click: typing.Optional[str] = None
    week_no_vip_click: typing.Optional[str] = None
    total_recommend: typing.Optional[str] = None
    month_recommend: typing.Optional[str] = None
    week_recommend: typing.Optional[str] = None
    total_favor: typing.Optional[str] = None
    month_favor: typing.Optional[str] = None
    week_favor: typing.Optional[str] = None
    current_yp: typing.Optional[str] = None
    total_yp: typing.Optional[str] = None
    current_blade: typing.Optional[str] = None
    total_blade: typing.Optional[str] = None
    week_fans_value: typing.Optional[str] = None
    month_fans_value: typing.Optional[str] = None
    total_fans_value: typing.Optional[str] = None
    last_chapter_info: LastChapterInfo
    tag_list: typing.List[TagList]
    book_type: typing.Optional[str] = None
    transverse_cover: typing.Optional[str] = None


class ChapterInfo(pydantic.BaseModel):
    chapter_id: typing.Optional[str] = None
    book_id: typing.Optional[str] = None
    division_id: typing.Optional[str] = None
    unit_hlb: typing.Optional[str] = None
    chapter_index: typing.Optional[str] = None
    chapter_title: typing.Optional[str] = None
    author_say: typing.Optional[str] = None
    word_count: typing.Optional[str] = None
    discount: typing.Optional[str] = None
    is_paid: typing.Optional[str] = None
    auth_access: typing.Optional[str] = None
    buy_amount: typing.Optional[str] = None
    tsukkomi_amount: typing.Optional[str] = None
    total_hlb: typing.Optional[str] = None
    uptime: typing.Optional[str] = None
    mtime: typing.Optional[str] = None
    ctime: typing.Optional[str] = None
    base_status: typing.Optional[str] = None
    txt_content: typing.Optional[str] = None


class ContentInfo(pydantic.BaseModel):
    chapter_id: typing.Optional[str] = None
    chapter_index: typing.Optional[str] = None
    chapter_title: typing.Optional[str] = None
    word_count: typing.Optional[str] = None
    tsukkomi_amount: typing.Optional[str] = None
    is_paid: typing.Optional[str] = None
    mtime: typing.Optional[str] = None
    is_valid: typing.Optional[str] = None
    auth_access: typing.Optional[str] = None
    text_content_path: typing.Optional[str] = None


class DivisionList(pydantic.BaseModel):
    max_update_time: typing.Optional[str] = None
    max_chapter_index: typing.Optional[str] = None
    division_id: typing.Optional[str] = None
    division_index: typing.Optional[str] = None
    division_name: typing.Optional[str] = None
    chapter_list: typing.List[ContentInfo]
