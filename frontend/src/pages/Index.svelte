<script>
    import { Link } from "svelte-routing";
    import loginHeaderSVG from "assets/loginHeader.svg";
    import googleSVG from "assets/google.svg";
    import facebookSVG from "assets/facebook.svg";
    import Input from "components/Input";
    import auth from "utils/auth";

    let username = "",
        usernameError = "";
    let password = "",
        passwordError = "";

    /* async function doGoogle() { */
    /*     fetch(API_URL + "/auth/google", { redirect: "follow" }); */
    /* } */

    async function doLogin() {
        let res = await fetch(API_URL + "/user/login", {
            method: "POST",
            body: JSON.stringify({
                username,
                password,
            }),
            credentials: "include",
        });

        usernameError = "";
        passwordError = "";

        switch (res.status) {
            case 404:
                usernameError = "User not found";
                return;
            case 401:
                passwordError = "Pasword incorrect";
                return;
        }

        $auth = (await res.json()).username;
    }
</script>

<section>
    {@html loginHeaderSVG}
    <form on:submit|preventDefault={doLogin}>
        <header>
            <h1>Enter to Brilla!</h1>
            <!-- <div>
                <a href="/api/auth/google" class="google">
                    <span>Login with </span>{@html googleSVG}
                </a>
                <a href="/api/auth/facebook" class="facebook">
                    <span>Login with </span>{@html facebookSVG}
                </a>
            </div> -->
        </header>
        <Input
            type="text"
            label="User"
            id="username"
            bind:value={username}
            errorMessage={usernameError}
        />

        <Input
            type="password"
            label="Password"
            id="password"
            style="-webkit-text-security: '⭐';"
            bind:value={password}
            errorMessage={passwordError}
        />

        <div>
            <Link to="/signup">Sign up</Link>
            <input type="submit" value="Log in" />
        </div>
    </form>
    <div />
</section>

<style lang="scss">
    section {
        display: flex;
        flex-direction: column;
        height: 100vh;
        justify-content: space-between;

        > :global(svg) {
            align-self: flex-start;
            order: from-md(99);
            transform: from-md(scaleY(-1));
            max-height: from-md(50vw);
        }

        form {
            display: grid;
            padding: 32px;
            grid-gap: 16px;
            width: from-md(60%) from-lg(50%);
            margin-left: auto;

            header {
                display: flex;
                justify-content: space-between;
                align-items: flex-end;
                > div {
                    display: flex;
                    align-items: end;
                    button,
                    a {
                        --background: white;
                        --shadow-color: white;

                        display: flex;
                        align-items: center;
                        border: none;
                        background: var(--background);
                        size: 48px from-md(auto 48px);
                        margin-left: 8px;
                        border-radius: 8px;
                        padding: 8px;
                        box-shadow: 0px 4px 8px var(--shadow-color);
                        span {
                            display: none from-md(initial);
                            font-weight: bold;
                            color: white;
                            margin: 0 8px 0 0;
                            font-size: 1.25rem;
                        }
                        &.google {
                            --background: linear-gradient(
                                45deg,
                                #db4437,
                                #dc5429
                            );
                            --shadow-color: rgba(219, 68, 55, 0.3);
                        }

                        &.facebook {
                            --background: linear-gradient(
                                45deg,
                                #4267b2,
                                #4254b2
                            );
                            --shadow-color: rgba(66, 103, 178, 0.3);
                        }
                        :global(svg) {
                            size: auto 100%;
                            fill: white;
                        }
                    }
                }
            }

            > div {
                display: flex;
                justify-content: space-between;

                :global(a),
                input {
                    padding: 12px 32px;
                    text-align: center;
                    border-radius: 12px;
                    font-weight: bold;
                    font-size: 1rem;
                    cursor: pointer;
                }

                :global(a) {
                    color: var(--secondary-color);
                    /* border: 2px solid var(--secondary-color); */
                    /* box-shadow: 0px 4px 8px rgba(251, 86, 102, 0.3); */
                }

                input {
                    background-color: var(--primary-color);
                    border: none;
                    box-shadow: 0 4px 8px rgba(249, 197, 95, 0.3);
                }
            }
        }

        > :global(a) {
            color: white;
            font-weight: bold;
            text-align: center;
            padding: 16px;
        }
    }
</style>
