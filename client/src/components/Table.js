import MUIDataTable from "mui-datatables";
import React, { useEffect, useState } from "react";
import TableRow from "@material-ui/core/TableRow";
import TableCell from "@material-ui/core/TableCell";
import ButtonGroup from '@material-ui/core/ButtonGroup';
import RatingButtonGroup from './Rating'
import CopyToClipboardButtonGroup from './CopyToClipboard'
import DeleteButton from './Delete'
import Logo from './acumen.png'

function Table(props) {

  const [rows, setRows] = useState([]);
  const [clipboardTemplates, setClipboardTemplates] = useState([])

  useEffect(()=>{
    fetch(props.api_url)
      .then(res => res.json())
      .then(
        (response) => {
          setRows(response.results);
          setClipboardTemplates(response.clipboard_templates)
        })
  }, [props.api_url])

  const updateRow = (rowId, obj) => {
    setRows([
        ...rows.slice(0,rowId).concat(
          Object.assign({}, rows[rowId], obj)
        ),
        ...rows.slice(rowId+1)
      ]);
  }

  const updateRating = (rowId, rating) => {
    updateRow(rowId, {rating: rating})
  }

  const deleteRow = (rowId) => {
    setRows(rows.filter((row, index) => index !== rowId));
  }

  const options = {
    print: false,
    responsive: 'standard',
    draggableColumns: {
      enabled: true,
    },
    filterType: 'multiselect',

    expandableRows: true,
    expandableRowsHeader: false,
    expandableRowsOnClick: true,
    renderExpandableRow: (rowData, rowMeta) => {
      const colSpan = rowData.length + 1;
      return (
        <TableRow>
          <TableCell colSpan={colSpan}>
            <ButtonGroup orientation="vertical">
              <RatingButtonGroup updateRating={updateRating} index={rowMeta.dataIndex}></RatingButtonGroup>
              <CopyToClipboardButtonGroup
                rowData = {rowData}
                buttons = {clipboardTemplates}
              ></CopyToClipboardButtonGroup>
              <DeleteButton
                deleteRow = {deleteRow}
                index = {rowMeta.dataIndex}
              ></DeleteButton>
            </ButtonGroup>
          </TableCell>
        </TableRow>
      );
    },
  };

  return (
    <div>
      <MUIDataTable
        title={
          <div style={{
            float: "left",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
          }}>
            <img src={Logo} height="100px"></img>
            <h2>
              {props.title}
            </h2>
          </div>
        }
        data={rows}
        columns={props.columns}
        options={options}
      />
    </div>
  )
}

export default Table;
