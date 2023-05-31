-- TODO: answer here
SELECT reports.id, students.fullname, students.class, students.status, reports.study, reports.score
FROM reports
JOIN students ON reports.student_id = students.id
WHERE students.status = 'active' AND reports.score < 70
ORDER BY reports.score;