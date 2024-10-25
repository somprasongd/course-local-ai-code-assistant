# @Postgresql

## Prompt

```text
@public.t_patient @public.f_patient_prefix
create select sql to find patient fullname with format:

<public.f_patient_prefix.patient_prefix_description><public.t_patient.patient_firstname> <public.t_patient.patient_lastname>
```

## Output

To generate a SQL query that joins the `t_patient` table with the `f_patient_prefix` table and formats the full name based on the given requirements, you can use the following SQL:

```sql
SELECT 
    CONCAT(
        COALESCE(fpp.patient_prefix_description_eng, fpp.patient_prefix_description),
        ' ',
        tp.patient_firstname,
        ' ',
        tp.patient_lastname
    ) AS patient_fullname
FROM 
    public.t_patient tp
JOIN 
    public.f_patient_prefix fpp ON tp.f_patient_prefix_id = fpp.f_patient_prefix_id;
```

### Explanation

1. **SELECT Clause**:
   - `CONCAT`: This function is used to concatenate strings together.
   - `COALESCE(fpp.patient_prefix_description_eng, fpp.patient_prefix_description)`: The `COALESCE` function returns the first non-null value in the list. If `patient_prefix_description_eng` is not null, it will be used; otherwise, `patient_prefix_description` will be used.

2. **FROM Clause**:
   - Specifies the main table from which to select data (`public.t_patient tp`).

3. **JOIN Clause**:
   - Joins the `t_patient` table with the `f_patient_prefix` table on the column `f_patient_prefix_id`.

### Result

The result will be a list of full names formatted as specified: `<patient_prefix_description> <first_name> <last_name>`.

For example, if there is an entry in `f_patient_prefix` like:

```sql
{
    "f_patient_prefix_id": "003",
    "patient_prefix_description": "นาย",
    "patient_prefix_description_eng": "Mr."
}
```

And in `t_patient`:

```sql
{
    "f_patient_prefix_id": "003",
    "patient_firstname": "พิศสมัย",
    "patient_lastname": "นามณี"
}
```

The query will return: `Mr. พิศสมัย นามณี`.
