# %%
import sqlalchemy
import pandas as pd
engine = sqlalchemy.create_engine('postgresql://postgres:postgres@db:5432/captrivia')

# %%
results = pd.read_sql('SELECT * FROM daily ORDER BY start_date', engine)
results

# %%
from matplotlib import pyplot as plt 
results.plot(x='start_date', rot=45, xticks=results['start_date'])

# %%
