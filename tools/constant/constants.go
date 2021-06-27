package constants

/**
 * 注册登录 platform 定值：open
 */
const defaultPlatform = "open"

const ApiV1 = "v1"
const ApiV2 = "v2"

/**
 * @Description: 登录相关
 */
const UserRegisterRequestUri = "open/register" //用户
const UserLoginRequestUri = "open/login"

/**
 * @Description: 会议相关
 */
const LiveListRequestUri = "open/room"                     //会议列表
const LiveCreateRequestUri = "open/room"                   //创建会议
const LiveUpdateRequestUri = "open/room/%d"                //会议修改
const LiveDetailRequestUri = "open/room/%d"                //会议详情
const LiveStatusRequestUri = "open/room/%d/status"         //修改会议状态
const LiveAccessRequestUri = "open/room/%d/access"         //用户提前入会
const LiveRecordRequestUri = "open/room/%d/record"         //会议录播列表
const LiveScreenShotRequestUri = "open/room/%d/screenshot" //会议截图列表
const LiveAddSpeakerRequestUri = "open/room/%d/speaker"    //会议增加讲者

/**
 * @Description: 会议飞检相关
 */
const LiveFlyCheckRequestUri = "open/room/check"      //会议飞检接口
const LiveFlyLayoutRequestUri = "open/room/%d/layout" //获取房间布局信息
const LiveFlyMemberRequestUri = "open/room/%d/member" //获取在线列表

/**
 * @Description: 会议讨论接口
 */
const LiveChatListRequestUri = "open/room/%d/chats"     //讨论列表
const LiveChatCheckRequestUri = "open/room/%d/chats/%d" //讨论审核

/*
 * @Description: 文件相关接口
 */
const LiveFileUploadRequestUri = "open/file"             //文件上传
const LiveFileListRequestUri = "open/file"               // 文件列表
const LiveFileDeleteRequestUri = "open/file/%d"          // 文件删除
const LiveFilePrintListRequestUri = "open/file/%d/print" // 文件打印
const LiveFileBindRequestUri = "open/room/%d/file"       // 会议文件绑定
const LiveFileBindListRequestUri = "open/room/%d/file"   // 会议文件绑定列表
const LiveFileBindDeleteRequestUri = "open/room/%d/file" // 会议文件绑定删除

/**
 * @Description: 调研相关
 */
const SurveyCreateRequestUri = "open/survey"           //添加调研
const SurveyDetailRequestUri = "open/survey/%d"        //调研详情
const SurveyListRequestUri = "open/survey"             //调研列表
const SurveyUpdateRequestUri = "open/survey/%d"        //调研更新
const SurveyDeleteRequestUri = "open/survey/%d"        //删除调研
const SurveyStatusRequestUri = "open/survey/%d/status" //调研发布状态修改

/**
 * @Description:  调研问题相关
 */
const SurveyQuestionCreateRequestUri = "open/survey/%d/question"    //调研问题添加
const SurveyQuestionUpdateRequestUri = "open/survey/%d/question/%d" //调研问题更新
const SurveyQuestionDeleteRequestUri = "open/survey/%d/question/%d" //调研问题删除
const SurveyQuestionListRequestUri = "open/survey/%d/question"      //调研问题列表
const SurveyQuestionDetailRequestUri = "open/survey/%d/question/%d" //调研问题详情

/**
 * @Description: 菜单相关
 */
const MenuCreateRequestUri = "open/room/%d/menu"
const MenuUpdateRequestUri = "open/room/%d/menu/%s"
const MenuListRequestUri = "open/room/%d/menu"
const MenuDeleteRequestUri = "open/room/%d/menu/%s"
const MenuSortRequestUri = "open/room/%d/menu_sort"

/**
 * @Description: 数据
 */
const DataLiveBaseRequestUri = "open/room/%d/live_base"                                           //会议基础数据统计
const DataUserJoinLogsUri = "open/room/%d/user_join_logs"                                         //用户出入记录
const DataUserActionLogsUri = "open/room/%d/user_action_logs"                                     //用户行为记录
const DataAccessRecordsRequestUri = "open/room/%d/access_records"                                 //会议出入记录
const DataSurveyRecordsRequestUri = "open/room/%d/survey_records"                                 //会议调研统计
const DataAboutLiveRequestUri = "open/survey/%d/about_live"                                       //会议调研报表
const DataSurveyInfosRequestUri = "open/room/%d/survey_infos"                                     //调研参与记录
const DataRoomLogsRequestUri = "open/room/%d/room_logs"                                           //会议用户出入记录
const DataWebsocketChatsRequestUri = "open/room/%d/websocket_chats"                               //会议留言导出
const DataUserAllLogsRequestUri = "open/room/%d/user_all_logs"                                    //用户参会总时长
const DataUserVodJoinLogsRequestUri = "open/room/%d/user_vod_join_logs"                           //录播观看出入记录
const DataLiveReportRequestUri = "open/room/%d/live_report"                                       //用户参会有效时长
const DataVodReportRequestUri = "open/room/%d/vod_report"                                         //用户录播观看时长
const DataOnlineRecordLogsRequestUri = "open/room/%d/online_record_logs"                          //参会人数统计日志
const DataUserLiveRequestUri = "open/room/%d/user_live"                                           //会议用户列表
const DataUserLiveAllLogsRequestUri = "open/room/%d/user_live_all_logs"                           //用户直播观看总时长统计
const DataUserLiveLogsByRtcTimeRequestUri = "open/room/%d/user_live_logs_by_rtc_time"             //根据讲者上下线时间做用户观看时长统计
const DataUserLiveUniqueLogsByRtcTimeRequestUri = "open/room/%d/user_live_uniquelogs_by_rtc_time" //根据讲者上下线时间做用户观看时长统计并合并用户终端
