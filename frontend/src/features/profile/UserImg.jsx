import Card from "../../components/card/Card"

export default function UserImg(props) {
  return (
    <Card styleName={props.userImgHolder}>
        <img className={props.profileImg} src={props.src} />
    </Card>
  )
}
