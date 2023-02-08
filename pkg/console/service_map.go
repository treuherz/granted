package console

// ServiceMap maps CLI flags to AWS console URL paths.
// e.g. passing in `-s ec2` will open the console at the ec2/v2 URL.
var ServiceMap = map[string]string{
	"":               "console",
	"athena":         "athena",
	"appsync":        "appsync",
	"c9":             "cloud9",
	"ce":             "cost-management",
	"cf":             "cloudfront",
	"cfn":            "cloudformation",
	"cloudformation": "cloudformation",
	"cloudmap":       "cloudmap",
	"cloudwatch":     "cloudwatch",
	"config":         "config",
	"ct":             "cloudtrail",
	"cw":             "cloudwatch",
	"ddb":            "dynamodbv2",
	"dms":            "dms/v2",
	"dx":             "directconnect/v2",
	"eb":             "elasticbeanstalk",
	"ebs":            "elasticbeanstalk",
	"ec2":            "ec2/v2",
	"ecr":            "ecr",
	"ecs":            "ecs",
	"eks":            "eks",
	"gd":             "guardduty",
	"grafana":        "grafana",
	"iam":            "iamv2",
	"l":              "lambda",
	"lambda":         "lambda",
	"mwaa":           "mwaa",
	"param":          "systems-manager/parameters",
	"r53":            "route53/v2",
	"rds":            "rds",
	"redshift":       "redshiftv2",
	"route53":        "route53/v2",
	"s3":             "s3",
	"sagemaker":      "sagemaker",
	"secretsmanager": "secretsmanager",
	"scrm":           "secretsmanager",
	"securityhub":    "securityhub",
	"sfn":			  "states",
	"scrh":           "securityhub",
	"ses":            "ses",
	"stepfn":         "states",
	"ssm":            "systems-manager",
	"sns":            "sns",
	"states":		  "states",
	"sso":            "singlesignon",
	"trustedadvisor": "trustedadvisor",
	"tra":            "trustedadvisor",
	"vpc":            "vpc",
	"waf":            "wafv2",
}

var globalServiceMap = map[string]bool{
	"dx":             true,
	"iam":            true,
	"r53":            true,
	"route53":        true,
	"trustedadvisor": true,
}
