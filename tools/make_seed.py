import json

with open('faculty.json', 'r', encoding='utf-8') as f:
    d = json.load(f)

output = ""

values = ', '.join([f"('{faculty['code']}','{faculty['name']}')" for faculty in d])
s = f"INSERT INTO faculties(code, name_th) VALUES {values};"
output += s

values = []
for faculty in d:
    for department in faculty['departments']:
        values.append(f"('{department['code']}', '{faculty['code']}', '{department['name']}')")
s = f"INSERT INTO departments(code, faculty_code, name_th) VALUES {', '.join(values)};"
output += s

values = []
for faculty in d:
    for department in faculty['departments']:
        for section in department['sections']:
            values.append(f"('{section['code']}', '{faculty['code']}', '{department['code']}', '{section['name']}')")
s = f"INSERT INTO sections(code, faculty_code, department_code, name_th) VALUES {', '.join(values)};"
output += s

with open('t.sql', 'w', encoding='utf-8') as f:
    f.write(output)
