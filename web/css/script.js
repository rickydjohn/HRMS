var countBox =2;

function addExp(){
	document.getElementById("countexp").value= countBox;
	var newChild = document.createElement("tr");
	document.getElementById("countexp").value= countBox;
	var row = document.getElementById("exp").insertRow(-1);
	var cell1 = row.insertCell(0);
	var cell2 = row.insertCell(1);
	var cell3 = row.insertCell(2);
	var cell4 = row.insertCell(3);
	var cell5 = row.insertCell(4);
	var cell6 = row.insertCell(5);
	cell1.innerHTML = "Company Name";
	cell2.innerHTML = "<input type='text' id='comp"+ countBox +"' name='comp"+ countBox+"'/>";  
	cell3.innerHTML = "From"; 
	cell4.innerHTML = "<input type='month' name='from"+ countBox +"'id='from"+countBox +"''/>";
	cell5.innerHTML = "To"; 
	cell6.innerHTML = "<input type='month' name='to"+ countBox +"'id='to"+countBox +"''/>";
	countBox += 1;
}

function addEdu(){
	document.getElementById("countedu").value= countBox;
	var newChild = document.createElement("tr");
	document.getElementById("countedu").value= countBox;
	var row = document.getElementById("edu").insertRow(-1);
	var cell1 = row.insertCell(0);
	var cell2 = row.insertCell(1);
	var cell3 = row.insertCell(2);
	var cell4 = row.insertCell(3);
	var cell5 = row.insertCell(4);
	var cell6 = row.insertCell(5);
	cell1.innerHTML = "Institution";
	cell2.innerHTML = "<input type='text' id='inst"+ countBox +"' name='inst"+ countBox+"'/>";  
	cell3.innerHTML = "Course"; 
	cell4.innerHTML = "<input type='text' name='course"+ countBox +"'id='course"+countBox +"'/>";
	cell5.innerHTML = "Passing Year"; 
	cell6.innerHTML = "<input type='month' name='yop"+ countBox +"'id='yop"+countBox +"'/>";
	countBox += 1;
}
