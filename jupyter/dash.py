# %%
import sys
sys.path.append('/Users/reh/.pyenv/versions/bot/lib/python3.11/site-packages')

# %%
import sqlalchemy
import pandas as pd
engine = sqlalchemy.create_engine('postgresql://postgres:postgres@localhost:5432/captrivia')

# %%
results = pd.read_sql('SELECT * FROM daily', engine)
results

# %%
from matplotlib import pyplot as plt 
results.plot(x='start_date', rot=45, xticks=results['start_date'])

# %%
