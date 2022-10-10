import Card from "../../components/card/Card"
export default function ProfilePosts(props) {
  return (
    <section className="profilePostSection">
    <Card styleName={props.contentSelector}>
            <span className={props.postBtn}><i className="fa-solid fa-table-list"></i> Posts</span>
            <span className={props.favBtn}><i className="fa-solid fa-bookmark"></i> Favourites</span>
            <span className={props.likeBtn}><i className="fa-solid fa-heart"></i> Liked</span>
    </Card>
    <Card styleName={props.postContainer}>
    </Card>
    </section>
  )
}
