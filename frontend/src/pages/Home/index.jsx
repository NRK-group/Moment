import './Home.css';

function Home() {
    return (
        <div className='homePage'>
            <h2>Home</h2>
            <div className='homePageContainer'>
                <div>
                    <div className='homePageStory'>Storys</div>
                    <div className='homePagePostArea'>
                        Post
                    </div>

                </div>
                <div className='homePageProfile'>homePageProfile</div>
            </div>
        </div>
    );
}

export default Home;
