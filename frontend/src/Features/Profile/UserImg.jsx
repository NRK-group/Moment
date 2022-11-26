import Card from "../../Components/Card/Card"

export default function UserImg(props) {
  return (
    <Card styleName={props.userImgHolder}>
        <img className={props.profileImg} src={props.src} />
    </Card>
  )
}
