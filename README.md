# jobcan-lambda

Simple [AWS Lambda](https://aws.amazon.com/lambda/) function to record your work attendance to [Jobcan](https://jobcan.ne.jp/).
This enables you to build a GPS-based automatic attendance recording system using [IFTTT](https://ifttt.com).

## Installation

- Create two AWS Lambda functions
  - jobcanStart
  - jobcanEnd
- Set env var `JOBCAN_AID` to `work_start` and `work_end` respectively
- Create [Amazon API Gateway](https://aws.amazon.com/api-gateway/) to trigger the functions
- Create IFTTT applets to call the API
  - Note that you have to send credential as body

## Usage

Below is an example body to send.

```json
{
  "credential": {
    "client_id": "your-company-name",
    "account_type": "staff",
    "login_id": "yourname@example.com",
    "password": "yoursecretpassword"
  }
}
```

## Note

Currently, once your API endpoint exposed, anyone around the world can invoke the function without any limitation.
You might want throttling and/or authentication to call the API.

## Future Works

- Detailed installation
- Guide to add authentication
- Publish IFTTT applet
