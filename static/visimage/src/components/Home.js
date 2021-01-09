import React, { Fragment } from "react";
import { useAuth0 } from "../react-auth0-spa";

const Home = () => {
    const { isAuthenticated, loginWithRedirect } = useAuth0();
    return (
        <Fragment>
            <div className="container">
                <div className="jumbotron text-center mt-5">
                    <h1>Visimage Library</h1>
                    <p>User image repository</p>
                    {!isAuthenticated && (
                        <button className="btn btn-primary btn-lg btn-login btn-block" onClick={() => loginWithRedirect({})}>Sign in</button>
                    )}
                </div>
            </div>
        </Fragment>
    );
};

export default Home;
