<script>
    import { onMount } from "svelte";
    import { Link } from "svelte-routing";
    import Popover from "svelte-popover";
    import debounce from "lodash/debounce";
    import Brillo from "components/Brillo";
    import Menu from "components/Menu";
    import FormBrillo from "components/FormCreateBrillo";
    import InfiniteLoading from "svelte-infinite-loading";
    import auth from "utils/auth";
    import { truncate, merge } from "utils/strings";
    import PencilPlus from "svelte-material-icons/PencilPlus";
    import Close from "svelte-material-icons/Close";
    import Magnify from "svelte-material-icons/Magnify";

    import Tumbleweed from "assets/tumbleweed.png";

    let brights = [];
    let userSearch = [];
    let brightSearch = [];
    let see = false;
    async function logout() {
        try {
            await fetch(API_URL + "/logout", { credentials: "include" });
            $auth = null;
        } catch (e) {}
    }

    async function fetchBrights(offset, event) {
        let data = await fetch(API_URL + `/timeline?offset=${offset}`);
        let jsonData = await data.json();
        jsonData.forEach((el) => brights.push(el));
        brights = brights;
        if (event != null && jsonData.length < 10) event.detail.complete();
        event?.detail.loaded();
    }

    const handleInput = debounce(async ({ target: { value: value } }) => {
        userSearch = [];
        let res = await fetch(
            API_URL + `/search?q=${encodeURIComponent(value)}`
        );
        let json = await res.json();
        json.forEach((el) => {
            let t = document.createElement("div");
            t.innerHTML = el.content;
            el.content = t.innerText;
        });
        console.log(json);
        if (value.startsWith("@")) userSearch = json;
        else brightSearch = json;
    }, 300);

    onMount(async () => {
        fetchBrights(0);
        // console.log(brights.length);
    });
</script>

<header>
    <Menu />
    <h1>Home</h1>
    <div>
        <span>
            <input on:input={handleInput} placeholder="Search" />
            <Magnify />
        </span>

        <div>
            <ul>
                {#each userSearch as user}
                    <li>
                        <Link to={`/user/${user.username}`}>
                            <div>
                                <h3>@{user.username}</h3>
                            </div>
                        </Link>
                    </li>
                {/each}
                {#each brightSearch as bright}
                    <li>
                        <Link to={`/brights/${bright["_key"]}`}>
                            <div>
                                <h3>@{bright.username}</h3>
                                <p>{truncate(bright.content, 50, true)}</p>
                            </div>
                        </Link>
                    </li>
                {/each}
            </ul>
        </div>
    </div>
</header>

<main>
    <section>
        {#if brights.length == 0}
            <p>No hay brillos para mostrar..</p>
        {:else}
            {#each brights as bright}
                <Brillo
                    user={{
                        username: bright.username,
                        name: bright.name,
                        img: `/media/${bright.userKey}/pp`,
                    }}
                    content={bright.content}
                    uploadDate={new Date(bright.created_at)}
                    id={bright._key}
                    rebrillos={bright.rebrillos}
                    interactions={bright.interactions}
                    ncomments={bright.ncomments}
                    comments={bright.comments}
                    media={bright.media}
                    rebrightedBy={bright.rebrightedBy}
                />
            {/each}
            <InfiniteLoading
                distance={200}
                on:infinite={(e) => fetchBrights(brights.length, e)}
            />
        {/if}
    </section>

    <div class:active={see}>
        <button on:click={() => (see = false)}><Close /></button>
        <FormBrillo route="/brights" />
    </div>

    <button on:click={() => (see = true)}><PencilPlus /></button>
</main>

<style lang="scss">
    main {
        display: grid;
        grid-gap: 16px;
        margin: 32px;
        > button {
            position: fixed;
            bottom: 16px;
            right: 16px;
            border: 2px solid var(--primary-color);
            background-color: var(--primary-color);

            border: 0;
            padding: 16px;
            border-radius: 100vmax;
            :global(svg) {
                display: block;
            }
        }

        section {
            height: 80vh;
            width: 100% from-md(80%) from-lg(60%);
            margin: 0 auto;
        }

        div {
            position: fixed;
            bottom: 0;
            transform: translateY(100%);
            text-align: end;

            &.active {
                transform: translateY(0);
                width: 100%;
                right: 0;
                background-color: var(--background-color);
                z-index: 999;
            }
            button {
                border: 0;
                background: var(--background-color);
                font-size: 16px;
                margin: 16px 24px 0 0;

                :global(svg) {
                    color: var(--dark-primary-color);
                    font-size: 24px;
                }
            }
        }

        :global(img) {
        }
    }

    header {
        display: grid;
        grid-template-columns: min-content 1fr min-content;
        justify-content: space-evenly;
        align-items: center;
        padding: 16px 24px from-md(16px 10%) from-lg(16px 20%);
        border-bottom: 1px solid var(--primary-color);
        box-shadow: 0px 4px 8px rgba(black, 0.2);

        h1 {
            font-size: 1.3rem;
        }

        > div {
            justify-self: end;
            align-self: end;
            > span {
                display: flex;
                align-items: center;
            }
            input {
                outline: none;
                border-radius: 16px;
                border: 1px solid black;
                padding: 4px 16px;
                &:hover,
                &:focus {
                    border-color: var(--secondary-color);
                }
            }
            > div {
                position: absolute;
                background-color: var(--background-color);
                z-index: 99;
            }
            p {
                font-size: 1.5em;
                font-weight: bolder;
            }
        }
        :global(button) {
            font-size: 1.5em;
            background: none;
            border: 0;
            padding: 8px 16px;
            align-self: center;
            height: 42px;
        }
    }

    ul {
        list-style: none;
        li {
            border: 1px solid var(--dark-primary-color);
            border-bottom: 0px;
            border-radius: 0px;
            padding: 8px;
            width: 100%;
            &:last-child {
                border-radius: 0px 0px 8px 8px;
                border: 1px solid var(--dark-primary-color);
            }
            &:first-child {
                border-radius: 8px 8px 0px 0px;
            }
        }
    }

    button {
        cursor: pointer;
    }

    @keyframes rebota {
        0% {
            transform: translate(0, 0), rotate(0deg);
        }
        30% {
            transform: translate(30vw, 0), rotate(360deg);
        }

        60% {
            transform: translate(50vw, 0), rotate(720deg);
        }
        100% {
            transform: translate(80vw, 0), rotate(1080deg);
        }
    }
</style>
