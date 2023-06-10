# IAM (Identity And Access Management)

Whenver we signup for AWS, a root user is created, this user access is used to create other roles and users in your organization or team. Users can also be clubbed into groups granting them same access levels.

- Users or Groups can be assigned to JSON documents called policies.
- These policies define the permisssions of users.
- In AWS, we apply the "least privilege prinnciple", i.e. do not give more permissions than a user needs.

## Creating an IAM user

- Go to IAM service from management console.
- Select **Users** from the left panel.
- Click on **Add user** button.
- For this example, we will create a user with Adimn-access, who has access to management console.
- Select `Provide access to AWS Management Console access`.
- Choose `I want to create an IAM user`.
- Select remaining options as per your requirement.
- Choose `Add user to group` and create a new group.
- If no group is created, then create a new group.
- For the sake of example, we will create a group with `AdminAccess` policy.
- This policy will give the user full access to AWS services.
- This is a managed policy, which is created by AWS.
- We can choose to add tags to the user. This is optional, but will help in organizing users.
- Once the user is created, and assigned to a group, we can either choose to email the credentials or download the creds.
- We can also choose to add the user to an existing group.

### Optional: Changing account Alias

- Go to IAM service from management console.
- Go to `Dashboard` from the left panel.
- On the right side, Under AWS Account, we can see Account ID and Account Alias.
- We can create a new alias by clicking on `Create` button.
- The alias can be used to login to AWS console.
