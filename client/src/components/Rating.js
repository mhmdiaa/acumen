import Button from '@material-ui/core/Button';
import ButtonGroup from '@material-ui/core/ButtonGroup';

function RatingButtonGroup(props) {
    const rating_limit = 10
    const buttons = []
  
    for (let i=1; i<=rating_limit; i++) {
      buttons.push(
        <Button onClick={() => {props.updateRating(props.index, i)}} key={i}>{i}</Button>
      )
    }

    return (
      <ButtonGroup color="primary" aria-label="outlined primary button group">
        {buttons}
      </ButtonGroup>
    )
}

export default RatingButtonGroup;
