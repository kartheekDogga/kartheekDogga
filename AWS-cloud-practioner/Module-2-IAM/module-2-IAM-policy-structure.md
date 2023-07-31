# IAM-Policy Structure

## JSON policy document structure

Each statement includes information about a single permission. If a policy includes multiple statements, AWS applies a logical OR across the statements when evaluating them. If multiple policies apply to a request, AWS applies a logical OR across all of those policies when evaluating them.

![image](https://docs.aws.amazon.com/images/IAM/latest/UserGuide/images/AccessPolicyLanguage_General_Policy_Structure.diagram.png)

The information in a statement is contained within a series of elements.

**Version** – Specify the version of the policy language that you want to use. We recommend that you use the latest 2012-10-17 version. For more information, see IAM JSON policy elements: Version

**Statement** – Use this main policy element as a container for the following elements. You can include more than one statement in a policy.

**Sid (Optional)** – Include an optional statement ID to differentiate between your statements.

**Effect** – Use Allow or Deny to indicate whether the policy allows or denies access.

**Principal (Required in only some circumstances)** – If you create a resource-based policy, you must indicate the account, user, role, or federated user to which you would like to allow or deny access. If you are creating an IAM permissions policy to attach to a user or role, you cannot include this element. The principal is implied as that user or role.

**Action** – Include a list of actions that the policy allows or denies.

**Resource (Required in only some circumstances)** – If you create an IAM permissions policy, you must specify a list of resources to which the actions apply. If you create a resource-based policy, this element is optional. If you do not include this element, then the resource to which the action applies is the resource to which the policy is attached.

**Condition (Optional)** – Specify the circumstances under which the policy grants permission.

To learn about these and other more advanced policy elements, see [IAM JSON policy elements reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html).

```{
  "Version": "2012-10-17", // another version is 2008-10-17, this is not recommended
  "Statement": [
    {
      "Sid": "FirstStatement", // optional
      "Effect": "Allow",
      "Action": ["iam:ChangePassword"],
      "Resource": "*"
    },
    {
      "Sid": "SecondStatement",
      "Effect": "Allow",
      "Action": "s3:ListAllMyBuckets",
      "Resource": "*"
    },
    {
      "Sid": "ThirdStatement",
      "Effect": "Allow",
      "Action": [
        "s3:List*",
        "s3:Get*"
      ],
      "Resource": [
        "arn:aws:s3:::confidential-data",
        "arn:aws:s3:::confidential-data/*"
      ],
      "Condition": {"Bool": {"aws:MultiFactorAuthPresent": "true"}} // optional but very useful
    }
  ]
}```

