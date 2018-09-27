package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type config struct {
	User string  `json:"user"`
	PWD  string  `json:"pwd"`
}

type populatedRestReference struct {
	Link  string `json:"link,omitempty"`
	Value string `json:"value,omitempty"`
}

type RestReference struct {
	populatedRestReference
}

type SBATasks struct {
	Result []struct {
		Parent                    RestReference `json:",omitonempty"`
		UPendingTriageChangedUser string        `json:"u_pending_triage_changed_user, omitempty"`
		WatchList                 string        `json:"watch_list, omitempty"`
		UAwrTemplate              string        `json:"u_awr_template, omitempty"`
		UponReject                string        `json:"upon_reject, omitempty"`
		SysUpdatedOn              string        `json:"sys_updated_on, omitempty"`
		UEscalationReason         string        `json:"u_escalation_reason, omitempty"`
		ApprovalHistory           string        `json:"approval_history, omitempty"`
		Skills                    string        `json:"skills, omitempty"`
		Number                    string        `json:"number, omitempty"`
		State                     string        `json:"state, omitempty"`
		USuppressAlerts           string        `json:"u_suppress_alerts, omitempty"`
		SysCreatedBy              string        `json:"sys_created_by, omitempty"`
		UCiBasedKbArticle         string        `json:"u_ci_based_kb_article, omitempty"`
		Knowledge                 string        `json:"knowledge, omitempty"`
		Order                     string        `json:"order, omitempty"`
		CmdbCi                    RestReference `json:"cmdb_ci"`
		Contract                  string        `json:"contract, omitempty"`
		Impact                    string        `json:"impact, omitempty"`
		Active                    string        `json:"active, omitempty"`
		WorkNotesList             string        `json:"work_notes_list, omitempty"`
		UVendor                   string        `json:"u_vendor, omitempty"`
		Priority                  string        `json:"priority, omitempty"`
		SysDomainPath             string        `json:"sys_domain_path, omitempty"`
		BusinessDuration          string        `json:"business_duration, omitempty"`
		GroupList                 string        `json:"group_list, omitempty"`
		UPendingTriageChangedDate string        `json:"u_pending_triage_changed_date, omitempty"`
		ApprovalSet               string        `json:"approval_set, omitempty"`
		UAffectedCiCount          string        `json:"u_affected_ci_count, omitempty"`
		ShortDescription          string        `json:"short_description"`
		UTaskPlannedEnd           string        `json:"u_task_planned_end, omitempty"`
		CorrelationDisplay        string        `json:"correlation_display, omitempty"`
		WorkStart                 string        `json:"work_start, omitempty"`
		AdditionalAssigneeList    string        `json:"additional_assignee_list, omitempty"`
		ULeadTime                 string        `json:"u_lead_time, omitempty"`
		ServiceOffering           string        `json:"service_offering, omitempty"`
		SysClassName              string        `json:"sys_class_name, omitempty"`
		ClosedBy                  RestReference `json:"closed_by, omitempty"`
		FollowUp                  string        `json:"follow_up, omitempty"`
		UNextUpdate               string        `json:"u_next_update, omitempty"`
		UMakeStub                 string        `json:"u_make_stub, omitempty"`
		ReassignmentCount         string        `json:"reassignment_count, omitempty"`
		AssignedTo                RestReference `json:"assigned_to, omitempty"`
		UResolvedB4Eos            string        `json:"u_resolved_b4_eos, omitempty"`
		UAwrAdhoc                 string        `json:"u_awr_adhoc, omitempty"`
		SLADue                    string        `json:"sla_due, omitempty"`
		CommentsAndWorkNotes      string        `json:"comments_and_work_notes, omitempty"`
		UAssignedBackOpen         string        `json:"u_assigned_back_open, omitempty"`
		Escalation                string        `json:"escalation, omitempty"`
		UTaskPlannedStart         string        `json:"u_task_planned_start, omitempty"`
		UponApproval              string        `json:"upon_approval, omitempty"`
		CorrelationID             string        `json:"correlation_id, omitempty"`
		UEventState               string        `json:"u_event_state, omitempty"`
		MadeSLA                   string        `json:"made_sla, omitempty"`
		UExtCorrelationID         string        `json:"u_ext_correlation_id, omitempty"`
		UHasAttachment            string        `json:"u_has_attachment, omitempty"`
		UWasATicket               string        `json:"u_was_a_ticket, omitempty"`
		UExtCorrelationDisplay    string        `json:"u_ext_correlation_display, omitempty"`
		UInactivityCount          string        `json:"u_inactivity_count, omitempty"`
		SysUpdatedBy              string        `json:"sys_updated_by, omitempty"`
		OpenedBy                  RestReference `json:"opened_by, omitempty"`
		USendApprovalNotif        string        `json:"u_send_approval_notif, omitempty"`
		UserInput                 string        `json:"user_input, omitempty"`
		SysCreatedOn              string        `json:"sys_created_on, omitempty"`
		SysDomain                 RestReference `json:"sys_domain, omitempty"`
		UAuthorizingContact       RestReference `json:"u_authorizing_contact, omitempty"`
		TaskFor                   RestReference `json:"task_for, omitempty"`
		ClosedAt                  string        `json:"closed_at, omitempty"`
		UOtherReason              string        `json:"u_other_reason, omitempty"`
		BusinessService           string        `json:"business_service, omitempty"`
		TimeWorked                string        `json:"time_worked, omitempty"`
		ExpectedStart             string        `json:"expected_start, omitempty"`
		OpenedAt                  string        `json:"opened_at"`
		USLABreachStatus          string        `json:"u_sla_breach_status, omitempty"`
		WorkEnd                   string        `json:"work_end"`
		WorkNotes                 string        `json:"work_notes, omitempty"`
		AssignmentGroup           RestReference `json:"assignment_group, omitempty"`
		Description               string        `json:"description, omitempty"`
		UVendorTicket             string        `json:"u_vendor_ticket, omitempty"`
		CalendarDuration          string        `json:"calendar_duration, omitempty"`
		CloseNotes                string        `json:"close_notes, omitempty"`
		UEmailBody                string        `json:"u_email_body, omitempty"`
		SysID                     string        `json:"sys_id, omitempty"`
		ContactType               string        `json:"contact_type, omitempty"`
		Urgency                   string        `json:"urgency, omitempty"`
		Company                   RestReference `json:"company, omitempty"`
		ActivityDue               string        `json:"activity_due, omitempty"`
		UDifficulty               string        `json:"u_difficulty, omitempty"`
		Comments                  string        `json:"comments, omitempty"`
		UVendorNote               string        `json:"u_vendor_note, omitempty"`
		Approval                  string        `json:"approval, omitempty"`
		DueDate                   string        `json:"due_date, omitempty"`
		SysModCount               string        `json:"sys_mod_count, omitempty"`
		SysTags                   string        `json:"sys_tags, omitempty"`
		URunTekSLA                string        `json:"u_run_tek_sla, omitempty"`
		UKbAttachment             string        `json:"u_kb_attachment, omitempty"`
		UEventGenerated           string        `json:"u_event_generated, omitempty"`
		Location                  string        `json:"location, omitempty"`
		UCustomerUpdate           string        `json:"u_customer_update, omitempty"`
	} `json:"result"`
}

func Unmarshal(data []byte, v interface{}) error {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string", vv)
	default:
		fmt.Println("Unknown type")
	}
	return nil
}

func printJSON(v interface{}) {
	switch vv := v.(type) {
	case string:
		fmt.Println("is string")
	case float64:
		fmt.Println("is float64", vv)
	case []interface{}:
		fmt.Println("is an array")
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	case map[string]interface{}:
		fmt.Println("is an object", vv)
		for i, u := range vv {
			fmt.Print(i, " ")
			printJSON(u)
		}
	default:
		fmt.Println("Unknown type")
	}
}

func LoadConfiguration(file string) config {
	var conf config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&conf)
	return conf
}

func main() {

	var buildTasks SBATasks
	cfile := "config.json"
	var f interface{}
	req, err := http.NewRequest("GET", "https://cdw.service-now.com/api/now/table/task?sysparm_query=company%3D967f64e86f5e4140174324dfae3ee498%5Eopened_at%3Ejavascript%3Ags.dateGenerate('2018-09-11'%2C'23%3A59%3A59')%5Eshort_description%3DAutomated%20Server%20Build%3A%20Create%20vm%20access%20script%20%26%20add%20entry%20to%20bmn%20%2Fetc%2Fhosts%5EORshort_description%3DApply%20the%20customer%20specific%20silo%20template%5EORshort_description%3DAutomated%20Server%20Build%3A%20Build%20Virtual%20Machine&sysparm_display_value=cmdb_ci&", nil)
	if err != nil {
		fmt.Println("Ma bad, must be a dagum typo.")
	}
	cStruct := LoadConfiguration(cfile)
	req.SetBasicAuth(cStruct.User, cStruct.PWD)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	bs, rerr := ioutil.ReadAll(resp.Body)
	if rerr != nil {
		// do something
	}

	json.Unmarshal(bs, &f)
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	printJSON((f))

	resp1, err1 := http.DefaultClient.Do(req)
	if err1 != nil {
		panic(err1.Error())
	}
	defer resp1.Body.Close()
	bs1, rerr1 := ioutil.ReadAll(resp1.Body)
	if rerr1 != nil {
		// do something
	}

	json.Unmarshal(bs1, &buildTasks)

}
