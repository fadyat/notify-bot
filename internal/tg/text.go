package tg

const (
	botDescription = `
<b>Hello, I'm a reminder bot!</b>
I can remind you about something you want to do.

<b>Message for creating a new reminder:</b>
<code>/new "name" date time frequency</code>

<b>Message for updating a reminder:</b>
<code>/update reminder_id "content" date time frequency</code>

<b>Message for getting all your reminders:</b>
<code>/get_all</code>

<b>Message for deleting a reminder:</b>
<code>/delete reminder_id</code>

<b>Date can be provided in the following formats:</b>
- <i>${number_of_days}</i>

<b>Frequency can be provided in the following formats:</b>
- <i>once</i>
- <i>daily</i>
- <i>weekly</i>
- <i>monthly</i>

<b>Example:</b>
<code>/new "Solve some Leetcode problems" 01-01 12:00 every_day</code>
`
)
