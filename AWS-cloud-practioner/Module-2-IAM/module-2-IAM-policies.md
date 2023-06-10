# IAM Poilicies

Policies are means in AWS to grant permissions to users, groups and roles. Policies are JSON documents that define the permissions. Policies can be attached to users, groups and roles.

IAM policies define permissions for an action regardless of the method that you use to perform the operation. For example, if a policy allows the GetUser action, then a user with that policy can get user information from the AWS Management Console, the AWS CLI, or the AWS API.

When you create an IAM user, you can choose to allow console(GUI) or programmatic access. If console access is allowed, the IAM user can sign in to the console(GUI) using their sign-in credentials. If programmatic access is allowed, the user can use access keys to work with the CLI or API.

## Identity-based policies

Identity-based policies are JSON permissions policy documents that control what actions an identity (users, groups of users, and roles) can perform, on which resources, and under what conditions. Identity-based policies can be further categorized:

- **Managed policies** – Standalone identity-based policies that you can attach to multiple users, groups, and roles in your AWS account. There are two types of managed policies:

  - **AWS managed policies** – Managed policies that are created and managed by AWS.
Customer managed policies – Managed policies that you create and manage in your AWS account. Customer managed policies provide more precise control over your policies than AWS managed policies.  

- **Inline policies** – Policies that you add directly to a single user, group, or role. Inline policies maintain ***a strict one-to-one relationship between a policy and an identity***. They are deleted when you delete the identity.

### Policy types

The following policy types, listed in order of usage-frequency, are available for use in AWS.

- **Identity-based policies** – Attach [managed](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#managedpolicy:~:text=be%20further%20categorized%3A-,Managed%20policies,-%E2%80%93%20Standalone%20identity%2Dbased) and [inline](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#managedpolicy:~:text=AWS%20managed%20policies.-,Inline%20policies,-%E2%80%93%20Policies%20that%20you) policies to IAM identities (users, groups to which users belong, or roles). Identity-based policies grant permissions to an identity.

- **Resource-based policies** – Attach inline policies to resources. The most common examples of resource-based policies are Amazon S3 bucket policies and IAM role trust policies. Resource-based policies grant permissions to the principal that is specified in the policy. Principals can be in the same account as the resource or in other accounts.

- **Permissions boundaries** – Use a managed policy as the permissions boundary for an IAM entity (user or role). That policy defines the maximum permissions that the identity-based policies can grant to an entity, but does not grant permissions. Permissions boundaries do not define the maximum permissions that a resource-based policy can grant to an entity.

- **Organizations SCPs** – Use an AWS Organizations service control policy (SCP) to define the maximum permissions for account members of an organization or organizational unit (OU). SCPs limit permissions that identity-based policies or resource-based policies grant to entities (users or roles) within the account, but do not grant permissions.

- **Access control lists (ACLs)** – Use ACLs to control which principals in other accounts can access the resource to which the ACL is attached. ACLs are similar to resource-based policies, although they are the only policy type that does not use the JSON policy document structure. ACLs are cross-account permissions policies that grant permissions to the specified principal. ACLs cannot grant permissions to entities within the same account.

- **Session policies** – Pass advanced session policies when you use the AWS CLI or AWS API to assume a role or a federated user. Session policies limit the permissions that the role or user's identity-based policies grant to the session. Session policies limit permissions for a created session, but do not grant permissions. For more information, see Session Policies.
