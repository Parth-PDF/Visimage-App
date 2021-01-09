import React, { useState, useEffect, useCallback } from "react";
import { useAuth0 } from "../react-auth0-spa";
import DeleteImage from "./DeleteImage";

const LoggedIn = () => {

    const [image, setImages] = useState([]);

    const {
        getTokenSilently,
        loading,
        user,
        logout,
        isAuthenticated,
    } = useAuth0();
    
    const getImages = useCallback(async () => {
        try {
            const token = await getTokenSilently();
            
            const response = await fetch("http://localhost:8080/images", {
                headers: {
                    Authorization: `Bearer ${token}`,
                },
            });

            const responseData = await response.json();

            setImages(responseData);
        } catch (error) {
            console.error(error);
        }
    }, [getTokenSilently])

    useEffect(() => {
        getImages();
    }, [getImages]);

    if (loading || !user) {
        return <div>Loading...</div>;
    }

    async function deletePicture(id) {
        
        const token = await getTokenSilently();

        const data = { id }
        await fetch("http://localhost:8080/delete",{
            headers: {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json",
            },
            method: 'DELETE',
            body: JSON.stringify(data),
        });

        getImages()
    }

    // CONVERT IMG FILE INTO BLOB AND HIT UPLOAD ENDPOINT TO ADD TO DB //
    function readFile(file) {
        return new Promise((resolve, reject) => {
            const fr = new FileReader();  
            fr.onload = function() {
                resolve(fr.result);
            };
            fr.readAsDataURL(file);
        });
    }

    async function uploadPicture(e) {
        const file = e.target.files[0];
        const datauri = await readFile(file);

        const token = await getTokenSilently();

        const data = { datauri };
        await fetch("http://localhost:8080/upload",{
            headers: {
                "Authorization": `Bearer ${token}`,
                "Content-Type": "application/json",
            },
            method: 'POST',
            body: JSON.stringify(data),
        });

        // get images after upload
        getImages()
    }

    return (
        <div className="container">
            {isAuthenticated && (
                <span
                    className="btn btn-primary float-right mt-3 mr-3"
                    onClick={() => logout()}
                >
                Log out
                </span>

            )}
            <div className="jumbotron text-center mt-5">
                <h1>Visimage Library</h1>

                <p>Welcome, {user.name}!</p>

                <div className="container">
                    <form>
                        <h4>Upload Image: <input onChange={uploadPicture} type="file" /></h4>
                    </form>
                </div>

                <br></br>

                <div className="row">
                    {image.map(function (image, index) {
                        var picture = image.imageTag
                        var tag = image.id
                        return (
                            <div className="col-sm-4" key={index}>
                                <div className="card mb-4">
                                    <div onClick={() => deletePicture(tag)} className="card-header">
                                        <DeleteImage />
                                    </div>
                                    <div className="card-body"> 
                                        <img src = {picture} alt = "Loading" width="200" height="300"/>
                                    </div>
                                    <div className="card-footer">
                                        Uploaded by {user.name}
                                    </div>
                                </div>
                            </div>
                        );
                    })}
                </div>
            </div>
        </div>
    );
};

export default LoggedIn;