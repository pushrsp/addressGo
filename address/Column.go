package address

var Columns = map[string][]string{
	"주소":   {"관리번호", "도로명코드", "읍면동일련번호", "지하여부", "건물본번", "건물부번", "기초구역번호", "변경사유코드", "고시일자", "변경전도로명주소", "상세주소부여여부"},
	"부가정보": {"관리번호", "행정동코드", "행정동명", "우편번호", "우편번호일련번호", "다량배달처명", "건축물대장건물명", "시군구건물명", "공동주택여부"},
	"지번":   {"관리번호", "일련번호", "법정동코드", "시도명", "시군구명", "법정읍면동명", "법정리명", "산여부", "지번본번", "지번부번", "대표여부", "건물본번", "건물부번", "읍면동명", "도로명", "도로명코드", "jibun_fullname", "doro_fullname"},
	"개선":   {"도로명코드", "도로명", "도로명로마자", "읍면동일련번호", "시도명", "시도로마자", "시군구명", "시군구로마자", "읍면동명", "읍면동로마자", "읍면동구분", "읍면동코드", "사용여부", "변경사유", "변경이력정보", "고시일자", "말소일자"},
}
