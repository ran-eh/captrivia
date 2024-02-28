FROM python:3.10-slim-buster

RUN apt-get update \
    && apt-get install -y --no-install-recommends

RUN apt-get update && apt-get install -y postgresql

WORKDIR /usr/src/dbt/dbt

# Install the dbt Postgres adapter. This step will also install dbt-core
RUN pip install --upgrade pip
RUN pip install dbt-postgres
RUN pip install pytz

# Install dbt dependencies (as specified in packages.yml file)
# Build seeds, models and snapshots (and run tests wherever applicable)
CMD PGPASSWORD=postgres psql -f ./analyses/up.pgsql -h db -U postgres -d captrivia && dbt run && sleep infinity

