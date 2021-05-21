import Button from '@material-ui/core/Button';
import DeleteIcon from '@material-ui/icons/Delete';

function DeleteButton(props) {
    return (
      <Button
        style={{
          textTransform: 'none',
          justifyContent: "flex-start"
        }}
        variant="outlined"
        color="secondary"
        size="large"
        startIcon={<DeleteIcon />}
        onClick = {() => {props.deleteRow(props.index)}}
      >
        Delete
      </Button>
    )
}

export default DeleteButton;
