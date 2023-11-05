package email

type Label string

const (
	// Email Types
	Correspondence Label = "correspondence"
	Solicitation   Label = "solicitation"
	Notification   Label = "notification"
	Alert          Label = "alert"
	Newsletter     Label = "newsletter"
	List           Label = "list"

	// Content Type
	Social            Label = "social"
	Personal          Label = "personal"
	Business          Label = "business"
	Professional      Label = "professional"
	Finance           Label = "finance"
	Bill              Label = "bill"
	Statement         Label = "statement"
	Receipt           Label = "receipt"
	OrderConfirmation Label = "order_confirmation"
	Shipping          Label = "shipping"
	Marketing         Label = "marketing"
	Sale              Label = "sale"
	Discount          Label = "discount"
	Health            Label = "health"
	Reminder          Label = "reminder"
	Education         Label = "education"
	Coursework        Label = "coursework"
	Travel            Label = "travel"
	Invite            Label = "invite"
	Event             Label = "event"
	News              Label = "news"
	Security          Label = "security"
	OneTimeLogin      Label = "one_time_login"
	ServiceUpdate     Label = "service_update"
	ToDo              Label = "todo"
	Reference         Label = "reference"
	Documents         Label = "documents"
	Spam              Label = "spam"
	Other             Label = "other"
)

var EmailTypes = map[Label]string{
	Correspondence: "An email from a sender of which you have responded too.",
	Solicitation:   "An email from an unknown sender that you have not engaged with.",
	Notification:   "A transactional email which is meant to notify you.",
	Alert:          "A transactional email which requires immediate attention.",
	Newsletter:     "A recurring email which is covers a given topic.",
	List:           "An email list",
}

var Content = map[Label]string{
	Social:            "An email which pertains to social networks or social events",
	Personal:          "An personal correspondance, contact, or communication.",
	Business:          "Emails related to commerce, trade, transactions, or other business activities.",
	Professional:      "Emails that pertain to your professional career or are work-related.",
	Finance:           "Emails dealing with monetary transactions, investments, or financial advice.",
	Bill:              "An email that contains a statement of money owed for goods or services.",
	Statement:         "An email that includes a detailed account statement, typically from financial institutions.",
	Receipt:           "An email that acknowledges the receipt of goods or services and the completion of a transaction.",
	OrderConfirmation: "An email confirming the details and acceptance of an order you have placed.",
	Shipping:          "An email providing information about the shipment of goods, including tracking details.",
	Marketing:         "Emails aimed at promoting products, services, or brands.",
	Sale:              "Emails that notify you of sales opportunities or special deals.",
	Discount:          "Emails offering reductions on regular prices of goods or services.",
	Health:            "Emails related to health services, medical information, or personal well-being.",
	Reminder:          "An email sent to prompt an action or recall an event, appointment, or task.",
	Education:         "Emails that contain educational content, learning resources, or academic information.",
	Coursework:        "Emails related to academic courses, assignments, or schoolwork.",
	Travel:            "Emails about travel arrangements, itineraries, and related promotions or information.",
	Invite:            "An email that invites you to an event, meeting, or social gathering.",
	Event:             "An email providing details about an upcoming event, including time, place, and agenda.",
	News:              "Emails that deliver news, updates on current affairs, or press releases.",
	Security:          "Emails concerning security alerts, privacy matters, or account protection.",
	OneTimeLogin:      "An email that provides a one-time login link or code for secure access.",
	ServiceUpdate:     "An email notifying you about updates, maintenance, or changes to a service you use.",
	ToDo:              "An email that lists tasks or actions that need to be completed.",
	Reference:         "An email that contains information or resources used for reference purposes.",
	Documents:         "An email that includes important documents or paperwork.",
	Spam:              "Unsolicited and often irrelevant emails, typically sent in bulk to a large number of users.",
	Other:             "An email that does not fit into the standard categories and is labeled as miscellaneous.",
}

// Algo
// Any email that I have responded to, that is not a newsletter, becomes a correspondance.
// Any new email that comes in will match this description

// Pull n emails
// Find all emails I have responded too
// mark senders are "coorespondants"
// any email that is received from those senders are marked as correspondence
// Any email that is CC'ed/Bcc'ed To, as well are correspondents

// Anything left that's not a Newsletter, List, or Correspondence is a notification or alert of somekind.
// Triage those with AI to figure out the content.
