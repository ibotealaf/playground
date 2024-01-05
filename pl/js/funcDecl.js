function printRecords(recordIds){
	function getStudentById() {
		const studentData = recordIds.map(function findStudent(id) {
			return studentRecords.find(function findStudentId(record) {
				return record.id == id
			})
		})
		return studentData
	}
	
	const studentsData = getStudentById()

	const sortedStudentData = studentsData.sort(function (data1, data2){
		if (data1.name < data2.name) return -1
		if (data1.name > data2.name) return 1
		return 0
	})

	sortedStudentData.forEach(function printInfo(student) {
		console.log(`${student.name} (${student.id}): ${student.paid ? "Paid" : "Not Paid"}`)

	})


}

function paidStudentsToEnroll() {
	const paidStudentsArr = studentRecords.filter(function paidStudent(data) {
		return data.paid
	})

	const paidStudentsId = paidStudentsArr.map(function getIds(student){
		return student.id
	})
	
	const studentsToEnroll = [...paidStudentsId]
	currentEnrollment.forEach((id) => {
		if (!studentsToEnroll.includes(id)) {
			studentsToEnroll.push(id)
		}
	})

	return studentsToEnroll


}

function remindUnpaid(recordIds) {

	const recordStudents = recordIds.map(id => studentRecords.find(student => student.id == id))
	return recordStudents
		.filter(student => !student.paid)
		.map(student => student.id)
}


var currentEnrollment = [410, 105, 664, 375]

var studentRecords = [
	{id: 313, name: "Frank", paid: true},
	{id: 410, name: "Suzy", paid: true},
	{id: 710, name: "Brian", paid: false},
	{id: 105, name: "Henry", paid: false},
	{id: 502, name: "Mary", paid: true},
	{id: 664, name: "Bob", paid: false},
	{id: 250, name: "Peter", paid: true},
	{id: 375, name: "Sarah", paid: true},
	{id: 867, name: "Greg", paid: false}
]

printRecords(currentEnrollment)
console.log("---------")
currentEnrollment = paidStudentsToEnroll()
printRecords(currentEnrollment)
console.log("---------")
let unpaidEnrollment = remindUnpaid(currentEnrollment)
printRecords(unpaidEnrollment)
