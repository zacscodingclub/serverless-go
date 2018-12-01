# Serverless Go on AWS

## Chapter 6

1. AWS CLI
2. Versions

- Versions are immutable and lock settings in place for a lambda.
- When publishing a new version of your Lambda function, you should give it a significant and meaningful version name that allows you to track different changes made to your function through its development cycle.
- SemVer <major>.<minor>.<patch> e.g. 1.1.2, 3.0.0
  - Major update means lambda is not backwards compatible
  - Minor indicates new features or functionality
  - Patch if fixing bugs or issues in previous version
- Alias is a pointer to a specific version, allowing you to promote a function from one environment to another (such as staging to production). Aliases are mutable, unlike versions.
