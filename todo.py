import csv


with open('all.csv') as f:
    reader = csv.reader(f)
    for row in reader:
        # first column is the task, second column is the status
        print(row[0],'|||' ,row[1])


