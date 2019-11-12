console.log(data);
var pl1 = document.getElementById("in1").value;
var pl2 = document.getElementById("in2").value;
function set1() {
    document.getElementById("player1").innerHTML = "<strong>Player1:</strong> "+document.getElementById("in1").value;
    document.getElementById("name1").innerHTML = "<strong>Player1:</strong> "+document.getElementById("in1").value;
}

function set2() {
    document.getElementById("player2").innerHTML = "<strong>Player2:</strong> "+document.getElementById("in2").value;
    document.getElementById("name2").innerHTML = "<strong>Player1:</strong> "+document.getElementById("in1").value;
}