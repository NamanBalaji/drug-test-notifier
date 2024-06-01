package browser_automation

const (
	loginPageURL                   = "https://rtlmo.my.salesforce-sites.com/Participant_LoginPage"
	logoutPage                     = "https://pnap-sarph-3802.my.salesforce.com/secur/logout.jsp"
	programIdSelector              = "#j_id0\\:j_id2\\:ProgramID"
	usernameSelector               = "#j_id0\\:j_id2\\:UserName"
	passwordSelector               = "#j_id0\\:j_id2\\:passWord"
	loginButtonSelector            = `input[name="j_id0:j_id2:j_id9"]`
	iFrameURL                      = "https://pnap-sarph-3802--c.vf.force.com/apex/WelcomePage?autoMapValues=1&inline=1&core.apexpages.framework.ApexViewServlet.getInlinedContentRequest=1&sfdcIFrameOrigin=https%3A%2F%2Fpnap-sarph-3802.my.salesforce.com%2Fhome%2Fhome.jsp&sdfcIFrameOrigin=https%3A%2F%2Fpnap-sarph-3802.my.salesforce.com%2Fhome%2Fhome.jsp"
	billsSelector                  = "#j_id0\\:j_id1\\:j_id2\\:noBalance span"
	testStatusButtonSelector       = `input[name="j_id0:j_id1:j_id2:j_id21:bottom:j_id22"]`
	selectionStatusSelectorYes     = "#j_id0\\:j_id1\\:j_id4 span"
	selectionStatusSelectorNo      = "#j_id0\\:j_id1\\:showRL span span:nth-of-type(3)"
	confirmationNumberSpanSelector = "#j_id0\\:j_id1\\:j_id14"
)
