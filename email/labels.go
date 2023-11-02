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
