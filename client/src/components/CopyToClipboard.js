import Button from '@material-ui/core/Button';
import ButtonGroup from '@material-ui/core/ButtonGroup';
import FileCopyIcon from '@material-ui/icons/FileCopy';

function copyToClipboard(text) {
navigator.clipboard.writeText(text).then(function() {
    // TODO: Display some kind of UI reaction?
}, function(err) {
    // TODO: Fallback or display a "browser not supported" error
});
}

function handleCopyToClipboardClick(template, data) {
    let rendered_template

    // Extremely hacky
    // Allows JS code in the template, so you can do basic text manipulation
    // Untrusted templates should never be used obviously
    try {
        rendered_template = eval("`" + template +"`")
    } catch(error) {
        // Should display a visible error here instead of this surprise
        rendered_template = error
    }
    copyToClipboard(rendered_template)
}

function CopyToClipboardButtonGroup(props) {
    let buttons = []

    let data = props.rowData
    let buttonDefinitions = props.buttons

    buttons = buttonDefinitions.map(
        (b, i) => {
            return (
                <Button
                onClick={() => { handleCopyToClipboardClick(b.template, data) }}
                style={{
                    textTransform: 'none',
                    justifyContent: "flex-start"
                }}
                key={i}
                startIcon={<FileCopyIcon />}
                >
                {b.label}
                </Button>
            )
        }
    )

    return (
        <ButtonGroup
            orientation="vertical"
            variant="outlined"
            color="primary"
            size="large"
        >
            {buttons}
        </ButtonGroup>
    )
}

export default CopyToClipboardButtonGroup;