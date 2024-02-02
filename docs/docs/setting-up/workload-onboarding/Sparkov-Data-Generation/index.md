---
sidebar_label: "Sparkov Data Generation"
sidebar_position: 11
---
# Generate Synthetic Data using Sparkov Data Generation technique


## Introduction

A synthetic dataset is generated by algorithms or simulations which has similar characteristics to real-world data. Collecting real-world data, especially data that contains sensitive user data like credit card information, is not possible due to security and privacy concerns. If a data scientist needs to train a model to detect credit fraud, they can use synthetically generated data instead of using real data without compromising the privacy of users.

The advantage of using Bacalhau is that you can generate terabytes of synthetic data without having to install any dependencies or store the data locally.

In this example, we will learn how to run Bacalhau on a synthetic dataset. We will generate synthetic credit card transaction data using the Sparkov program and store the results in IPFS.

### Prerequisite

To get started, you need to install the Bacalhau client, see more information [here](../../../getting-started/installation.md)

## 1. Running Sparkov Locally​

To run Sparkov locally, you'll need to clone the repo and install dependencies:



```bash
%%bash
git clone https://github.com/js-ts/Sparkov_Data_Generation/
pip3 install -r Sparkov_Data_Generation/requirements.txt
```

Go to the `Sparkov_Data_Generation` directory:


```python
%cd Sparkov_Data_Generation
```

Create a temporary directory (`outputs`) to store the outputs:


```bash
%%bash
mkdir ../outputs
```

## 2. Running the script

```bash
%%bash
python3 datagen.py -n 1000 -o ../outputs "01-01-2022" "10-01-2022"
```

The command above executes the Python script `datagen.py`, passing the following arguments to it:

`-n 1000`:  Number of customers to generate

`-o ../outputs`: path to store the outputs

`"01-01-2022"`: Start date

`"10-01-2022"`: End date

Thus, this command uses a Python script to generate synthetic credit card transaction data for the period from `01-01-2022` to `10-01-2022` and saves the results in the `../outputs` directory.


To see the full list of options, use:


```bash
%%bash
python datagen.py -h
```

## 3. Containerize Script using Docker

To build your own docker container, create a `Dockerfile`, which contains instructions to build your image:


```
%%writefile Dockerfile

FROM python:3.8

RUN apt update && apt install git

RUN git clone https://github.com/js-ts/Sparkov_Data_Generation/

WORKDIR /Sparkov_Data_Generation/

RUN pip3 install -r requirements.txt
```

These commands specify how the image will be built, and what extra requirements will be included. We use `python:3.8` as the base image, install `git`, clone the `Sparkov_Data_Generation` repository from GitHub, set the working directory inside the container to `/Sparkov_Data_Generation/`, and install Python dependencies listed in the `requirements.txt` file."

:::info
See more information on how to containerize your script/app [here](https://docs.docker.com/get-started/02_our_app/)
:::


### Build the container

We will run `docker build` command to build the container:

```
docker build -t <hub-user>/<repo-name>:<tag> .
```

Before running the command replace:

**`hub-user`** with your docker hub username. If you don’t have a docker hub account [follow these instructions to create docker account](https://docs.docker.com/docker-id/), and use the username of the account you created

**`repo-name`** with the name of the container, you can name it anything you want

**`tag`** this is not required but you can use the `latest` tag

In our case:

```
docker build -t jsacex/sparkov-data-generation .
```

### Push the container

Next, upload the image to the registry. This can be done by using the Docker hub username, repo name or tag.

```
docker push <hub-user>/<repo-name>:<tag>
```

In our case:

```
docker push jsacex/sparkov-data-generation
```


After the repo image has been pushed to Docker Hub, we can now use the container for running on Bacalhau

## 4. Running a Bacalhau Job


Now we're ready to run a Bacalhau job: 


```bash
%%bash --out job_id
bacalhau docker run \
    --id-only \
    --wait \
    jsacex/sparkov-data-generation \
    --  python3 datagen.py -n 1000 -o ../outputs "01-01-2022" "10-01-2022"
```

### Structure of the command:

`bacalhau docker run`: call to Bacalhau

`jsacex/sparkov-data-generation`: the name of the docker image we are using

`--  python3 datagen.py -n 1000 -o ../outputs "01-01-2022" "10-01-2022"`: the arguments passed into the container, specifying the execution of the Python script `datagen.py` with specific parameters, such as the amount of data, output path, and time range. 


When a job is submitted, Bacalhau prints out the related `job_id`. We store that in an environment variable so that we can reuse it later on:


```python
%env JOB_ID={job_id}
```

## 5. Checking the State of your Jobs

**Job status**: You can check the status of the job using `bacalhau list`.


```bash
%%bash
bacalhau list --id-filter ${JOB_ID}
```

When it says `Published` or `Completed`, that means the job is done, and we can get the results.

**Job information**: You can find out more information about your job by using `bacalhau describe`.



```bash
%%bash
bacalhau describe ${JOB_ID}
```

**Job download**: You can download your job results directly by using `bacalhau get`. Alternatively, you can choose to create a directory to store your results. In the command below, we created a directory (`results`) and downloaded our job output to be stored in that directory.


```bash
%%bash
rm -rf results && mkdir -p results
bacalhau get ${JOB_ID} --output-dir results
```


## 6. Viewing your Job Output

To view the contents of the current directory, run the following command:


```bash
%%bash
ls results/outputs  
```

## Support
If you have questions or need support or guidance, please reach out to the [Bacalhau team via Slack](https://bacalhauproject.slack.com/ssb/redirect) (**#general** channel).