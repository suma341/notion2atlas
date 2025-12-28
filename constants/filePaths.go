package constants

const NT_DATA_DIR = "notion_data"

const (
	CURRICULUM_DIR = NT_DATA_DIR + "/curriculums"
	PAGE_DIR       = NT_DATA_DIR + "/pages"
	CATEGORY_DIR   = NT_DATA_DIR + "/categories"
	INFO_DIR       = NT_DATA_DIR + "/infos"
	ANSWER_DIR     = NT_DATA_DIR + "/answers"
	OGP_DIR        = NT_DATA_DIR + "/ogp"
	ASSETS_DIR     = NT_DATA_DIR + "/assets"
	PAGE_DATA_DIR  = NT_DATA_DIR + "/pageData"
	TMP_DIR        = NT_DATA_DIR + "/tmp"
	SYNCED_DIR     = NT_DATA_DIR + "/synced"
)

const (
	// CURRICULUM_PATH     = CURRICULUM_DIR + "/data.json"
	CATEGORY_PATH           = CATEGORY_DIR + "/data.json"
	INFO_PATH               = INFO_DIR + "/data.json"
	ANSWER_PATH             = ANSWER_DIR + "/data.json"
	SYNCED_PATH             = SYNCED_DIR + "/data.json"
	TMP_PAGE_PATH           = TMP_DIR + "/page.json"
	TMP_ALL_PAGE_PATH       = TMP_DIR + "/all_page.json"
	PAGE_DAT_PATH           = PAGE_DIR + "/page.dat"
	SYNCED_DAT_PATH         = SYNCED_DIR + "/synced.dat"
	TMP_ALL_SYNCED_PATH     = TMP_DIR + "/all_synced.json"
	CURRICULUM_DAT_PATH     = CURRICULUM_DIR + "/curriculum.dat"
	TMP_ALL_CURRICULUM_PATH = TMP_DIR + "/curriculum.json"
)

const (
	TEST_DIR        = NT_DATA_DIR + "/test"
	TEST_PREV_DIR   = TEST_DIR + "/prev"
	TEST_RESULT_DIR = TEST_DIR + "/result"

	TEST_PREV_PAGE_PATH         = TEST_PREV_DIR + "/page.json"
	TEST_RESULT_PAGE_PATH       = TEST_RESULT_DIR + "/page.json"
	TEST_PREV_CURRICULUM_PATH   = TEST_PREV_DIR + "/curriculum.json"
	TEST_RESULT_CURRICULUM_PATH = TEST_RESULT_DIR + "/curriculum.json"
)
