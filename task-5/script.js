document.getElementById("Avatar").style.display = "none"
        document.getElementById("Identicon").style.display = "none"
        document.getElementById("Avatar").style.display = "none"
        document.getElementById("Usertxt").style.display = "none"
        document.getElementById("Nametxt").style.display = "none"
        document.getElementById("Biotxt").style.display = "none"
        document.getElementById("Mail").style.display = "none"
        document.getElementById("Mailtxt").style.display = "none"

        function getdata() {
            var login = document.getElementById("usin").value;
            var query = `query get_details($login: String!) {
                                user(login: $login) {
                                    login
                                    name
                                    avatarUrl
                                    bio
                                    email
                                }
                            }`
            fetch('https://api.github.com/graphql', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Accept': 'application/json',
                    'Authorization': 'bearer #################################',/*Use your key here. Use the link in readme file if you don't have a key and want to try this project*/
                },
                body: JSON.stringify({
                    query,
                    variables: {
                        login
                    },
                })
            }).then(r => r.json()).then(data => {
                console.log(data)
                if (data.data.user == null) {
                    document.getElementById("Username").innerHTML = "<h2>User does not exist<h2>"
                    document.getElementById("Usertxt").style.display = "none"
                    document.getElementById("Name").style.display = "none"
                    document.getElementById("Nametxt").style.display = "none"
                    document.getElementById("Bio").style.display = "none"
                    document.getElementById("Biotxt").style.display = "none"
                    document.getElementById("Avatar").style.display = "none"
                    document.getElementById("Identicon").style.display = "none"
                } else {
                    document.getElementById("Avatar").style.display = "inline"
                    document.getElementById("Name").style.display = "inline"
                    document.getElementById("Bio").style.display = "inline"
                    document.getElementById("Identicon").style.display = "inline"
                    document.getElementById("Avatar").style.display = "inline"
                    document.getElementById("Usertxt").style.display = "inline"
                    document.getElementById("Nametxt").style.display = "inline"
                    document.getElementById("Biotxt").style.display = "inline"
                    document.getElementById("Avatar").src = data.data.user.avatarUrl;
                    document.getElementById("Username").innerHTML = data.data.user.login;
                    document.getElementById("Identicon").src = "https://identicons.github.com/" + document.getElementById("usin").value + ".png";
                    document.getElementById("Name").innerHTML = data.data.user.name
                    document.getElementById("Bio").innerHTML = data.data.user.bio
                    document.getElementById("Mail").style.display = "inline"
                    document.getElementById("Mailtxt").style.display = "inline"
                    document.getElementById("Mail").innerHTML = data.data.user.email;
                    
                }
            });
        }
