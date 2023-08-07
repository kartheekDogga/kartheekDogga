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

## IAM security tools

- **IAM credentials report**: This report lists all the users in the account and the status of their various credentials, such as passwords, access keys, and MFA devices. This report helps you to ensure that your users have passwords and MFA devices, and that the users' access keys have not been exposed. This report also contains last used information for access keys and passwords, which you can use to identify unused credentials that you can delete.

- IAM Dashoard -> credential report -> download csv

- **IAM Access Advisor**: This tool helps you to identify which of your resources users can access based on their IAM permissions. You can use this information to identify which permissions are unused or not used often. You can then refine your policies to remove unnecessary permissions.

- IAM Dashboard -> Select User -> Access Advisor

### Optional: Changing account Alias

- Go to IAM service from management console.
- Go to `Dashboard` from the left panel.
- On the right side, Under AWS Account, we can see Account ID and Account Alias.
- We can create a new alias by clicking on `Create` button.
- The alias can be used to login to AWS console.

## IAM Best Practices

- Do not use root account for day-to-day activities. Except for AWS setup and billing.

- Create individual IAM users for everyone accessing AWS. 1 physical user = 1 IAM user.

- Assign users to groups and assign permissions to groups.

- create a strong password policy.

- Use and enforce the use of MFA.

- Create and use roles for giving permissions to AWS services.

- Use Access Keys for programmatic access (CLI/SDK) and not for console access.

## IAM Shared Responsibility Model

In AWS, While AWS is responsible for...

- Infrastructure (Golbal Network Security)

- Configuration and vulnerability Analysis

- Compliance validation and patching

We are responsible for...

- Users, Groups, Roles, Policies Management and Monitoring.

- Enforcing MFA, Strong Passwords, Credentials Rotation.

- Use IAM tools to apply least privilege principle.

- Analyze access patterns and review permissions.
