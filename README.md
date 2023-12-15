# Billing API Task

## Create an environment that allows a user to add a specific payment method to a contact. Each contact can have multiple payment methods. This payment method will consist of Insurance or Cash pay (via cash, CC, or debit card). Make sure that this is specifically connected to the contact ID of the contact it was added to.

## Current structure - User/Client adds all Payers (insurance Companies) they accept inside of the system as a Payer. When Adding a Billing/Payment to a contact, the system fires Payer API to search all payers added in the system and then puts that Payer as the contacts Billing method.

## Here are all the fields that are being used for the Billing/Payment for a specific contact.

- Payment Type (Insurance or Cash)
- Status (Active or Inactive)
- Payer (Fires API to search for All Payers in the system)
- Billing order (1,2,3,4 - what order the payment method should be used)
- Condition related to(Employment, Auto Accident, Other Accident,
  Applicable
- Billing ID ( Insurnace ID provided by Insurnace company)
- Relationship to Insured (Self, Spouse, Child, Other)
- Insured Full Name
- Insured ID (ID assigned to person Insured)
- Date Of Birth of Insured
- Address of Insured
- City of Insured
- States of insured
- Zip code of Insured
- Insurance Start Date
- Insurance End Date
- Insurance Co-Pay (how much Customers pays on each visit)
- Insurance Dedutable (Total amount out of pocket for customer)

### Create all API, Post, Get, Update, Delete

### Implement an API to manage the task and store the relevant data in the database. Develop CRUD REST APIs to manage all billing for contacts
