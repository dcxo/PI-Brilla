<script>
    import Input from "components/Input";
    import { Link, navigate } from "svelte-routing";
    import SVGheader from "assets/loginHeader.svg";
    import { calculateAge } from "utils/date";
    import DatePicker from "svelte-calendar";

    let username = "",
        usernameError = "";
    let password = "",
        passwordError = "";
    let name = "",
        nameError = "";
    let email = "",
        emailError = "";

    async function signup() {
        let res = await fetch(API_URL + "/user", {
            method: "POST",
            body: new URLSearchParams({
                username,
                password,
                name,
                email,
            }),
        });

        switch (res.status) {
            case 400:
                birthdayError = " Error birthday format";
                return;

            case 409:
                usernameError = "Username already exists";
                return;
        }

        window.location.href = "/";
    }

    async function check_username() {
        usernameError = "";
        passwordError = "";

        let res = await fetch(API_URL + "/user/login", {
            method: "POST",
            body: JSON.stringify({
                username,
                password: "",
            }),
        });

        switch (res.status) {
            case 401:
                usernameError = "User already exist";
                return;
            case 404:
                // error usuario no existe, error correcto
                if (username == "") {
                    usernameError = "User is required";
                } else if (password == "") {
                    passwordError = "Password is required";
                } else {
                    window.location.hash = "#id2";
                }
                console.log(usernameError);
                console.log(passwordError);
                return;
        }
    }

    async function check_email() {
        emailError = "";
        let res = await fetch(API_URL + "/user/exits", {
            method: "POST",
            body: JSON.stringify({
                email,
            }),
        });

        // console.log(res.status);
        switch (res.status) {
            case 409:
                emailError = "Email already exist";
                return;
        }
        if (emailError == "" && email != "") {
            window.location.hash = "#id3";
        } else {
            emailError = "Email is required ";
        }
    }
</script>

<section>
    {@html SVGheader}

    <h2>Create Account</h2>
    <form on:submit|preventDefault={signup}>
        <div id="id1">
            <Input
                type="text"
                label="Username"
                id="username"
                bind:value={username}
                bind:errorMessage={usernameError}
            />

            <Input
                type="password"
                label="Password"
                id="password"
                required
                bind:value={password}
                bind:errorMessage={passwordError}
            />

            <Link to="/">Before</Link>
            <a href="#id2" on:click|preventDefault={check_username}>Next</a>
        </div>

        <div id="id2">
            <Input
                type="email"
                label="Email"
                id="email"
                bind:value={email}
                bind:errorMessage={emailError}
                regex={/(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))/}
                invalidInputMessage="incorrect email"
            />
            <a href="#id1">Before </a>
            <a href="#id3" on:click|preventDefault={check_email}>Next</a>
        </div>
        <div id="id3">
            <Input
                type="text"
                label="Name"
                id="name"
                bind:value={name}
                bind:errorMessage={nameError}
            />

            <!-- ^\d{4}-([0]\d|1[0-2])-([0-2]\d|3[01])$ -->
            <div>
                <a href="#id2">Before</a>
                <input type="submit" value="Sign up" />
            </div>
        </div>
    </form>
</section>

<style lang="scss">
    section {
        // padding: 32px;
        display: flex;
        flex-direction: column;
        min-height: 100vh;
        justify-content: space-between;
        h2 {
            padding: 0 32px;
            margin-top: 16px;
        }
        > :global(svg) {
            align-self: flex-start;
            order: from-md(99);
            transform: from-md(scaleY(-1));
            max-height: from-md(30vw);
        }

        > form {
            display: flex;
            height: 70vh;
            align-items: center;
            overflow: hidden;
            > div {
                scroll-snap-align: center;
                width: 100%;
                flex: 0 0 100%;
                // margin: 16px;
                padding: 16px;

                :global(div) {
                    margin-bottom: 16px;
                    // padding: 0 16px;
                }
                :global(a),
                input[type="submit"] {
                    padding: 8px 16px;
                    text-align: center;
                    border-radius: 12px;
                    font-weight: bold;
                    background-color: var(--primary-color);
                    border: 2px solid var(--primary-color);
                    font-size: 14px;
                }
            }
        }
    }
</style>
