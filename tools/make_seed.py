import json

null_code = "-"
null_name = "-"

with open('faculty.json', 'r', encoding='utf-8') as f:
    d = json.load(f)

# append null code
for faculty in d:
    faculty['departments'].append({"code": null_code, "name_th": null_name, "name_en": null_name, "sections": []})
    for department in faculty['departments']:
        department['sections'].append({"code": null_code, "name_th": null_name, "name_en": null_name})

output = ""

values = ', '.join([f"('{faculty['code']}','{faculty['name_th']}','{faculty['name_en']}')" for faculty in d])
s = f"INSERT INTO faculties(code, name_th, name_en) VALUES {values};"
output += s

values = []
for faculty in d:
    for department in faculty['departments']:
        values.append(f"('{department['code']}', '{faculty['code']}', '{department['name_th']}', '{department['name_en']}')")
s = f"INSERT INTO departments(code, faculty_code, name_th, name_en) VALUES {', '.join(values)};"
output += s

values = []
for faculty in d:
    for department in faculty['departments']:
        for section in department['sections']:
            values.append(f"('{section['code']}', '{faculty['code']}', '{department['code']}', '{section['name_th']}', '{section['name_en']}')")
s = f"INSERT INTO sections(code, faculty_code, department_code, name_th, name_en) VALUES {', '.join(values)};"
output += s

with open('t.sql', 'w', encoding='utf-8') as f:
    f.write(output)

with open('nullable_faculty.json', 'w', encoding='utf-8') as f:
    json.dump(d, f, ensure_ascii=False)
