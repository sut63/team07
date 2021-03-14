import { useState, useEffect } from 'react';
import * as React from 'react';
import {
  Content,
  ContentHeader,
} from '@backstage/core';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntCarCheckInOut } from '../../api/models/EntCarCheckInOut';
import moment from 'moment';
import TextField from '@material-ui/core/TextField';
import { Alert } from '@material-ui/lab';
import SearchIcon from "@material-ui/icons/Search";
import DeleteForeverIcon from '@material-ui/icons/DeleteForever';

const useStyles = makeStyles({
 table: {
   minWidth: 1500,
 },
 formControl: {
  width: 400,
},
textField: {
  width: '25ch',
},
});

export default function ComponentsTable() {
 
 const classes = useStyles();
 const api = new DefaultApi();

 const [status, setStatus] = useState(false);
 const [loading, setLoading] = useState(true);
 const [alerttype, setAlertType] = useState(String);
 const [errormessege, setErrorMessege] = useState(String);
 const [search, setSearch] = useState(String);
 const [userid, setUser] = useState(Number);
 const [carcheckinout, setCarcheckinout] = useState<EntCarCheckInOut[]>([]);

 useEffect(() => {

  const checkJobPosition = async () => {
      const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
      setLoading(false);
      if (jobdata != "เจ้าหน้าที่รถพยาบาล" ) {
        localStorage.setItem("userdata",JSON.stringify(null));
        localStorage.setItem("jobpositiondata",JSON.stringify(null));
        history.pushState("","","./");
        window.location.reload(false);        
      }
      else {
        setUser(Number(localStorage.getItem("userdata")))
      }
    }
  checkJobPosition();

  const apiUrl = `http://localhost:8080/api/v1/carcheckinoutsearch?ambulance=${search}`;
  const requestOptions = {
    method: 'GET',
};  
  fetch(apiUrl, requestOptions)
    .then(response => response.json())
      .then(data => {
        console.log(data.data)
        setCarcheckinout([]);
        if (data.data != null) {
          if(data.data.length >= 1) {
            console.log("Data >= 1")
            setCarcheckinout(data.data);
        }
    }
    setStatus(true);
});

}, [loading]);  

const Searchcarinout = async () =>{
  const apiUrl = `http://localhost:8080/api/v1/carcheckinoutsearch?ambulance=${search}`;
  const requestOptions = {
    method: 'GET',
};  
  fetch(apiUrl, requestOptions)
    .then(response => response.json())
      .then(data => {
        console.log(data.data)
        setErrorMessege("ไม่พบข้อมูล");
        setAlertType("error");
        setCarcheckinout([]);
        if (data.data != null) {
          if(data.data.length >= 1) {
            console.log("Data >= 1")
            setErrorMessege("พบข้อมูล");
            setAlertType("success");
            setCarcheckinout(data.data);
        }
    }
    setStatus(true);
});
}

const SearchhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
  setSearch(event.target.value as string);
};

 const deleteCarcheckinouts = async (id: number) => {
    await api.deleteCarcheckinout({ id: id });
   setLoading(true);
 };

 return (
  <Content>
  <ContentHeader title="ค้นหารถเข้าออก">
    &nbsp;&nbsp;&nbsp;&nbsp;
    {status ? (
     <div>
      {alerttype != "" ? (
       <Alert severity={alerttype} onClose={() => { setStatus(false) }}>
          {errormessege}
       </Alert>
       ) : null}
     </div>
    ) : null}
  </ContentHeader>   
              <div>
            <TextField
                 id="search"
                 label = "ค้นหาทะเบียนรถ"
                 type="string"
                 value={search}
                 onChange={SearchhandleChange}
                 className={classes.textField}
                 style={{ width: 200 }}
                />
            &nbsp;&nbsp;&nbsp;&nbsp;
            <Button
              startIcon={<SearchIcon/>}
                onClick={() => {
                  Searchcarinout();
                }}
                variant="contained"
                color="primary"
              >
                ค้นหา
             </Button>
             </div>
          <br></br><br></br>
   
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">ทะเบียนรถ</TableCell>
           <TableCell align="center">เจ้าหน้าที่</TableCell>
           <TableCell align="center">วัตถุประสงค์</TableCell>
           <TableCell align="center">หมายเหตุ</TableCell>
           <TableCell align="center">จำนวนเจ้าหน้าที่</TableCell>
           <TableCell align="center">ระยะทาง</TableCell>
           <TableCell align="center">สถานที่</TableCell>
           <TableCell align="center">วันที่รถออก</TableCell>
           <TableCell align="center">วันที่รถเข้า</TableCell>
           <TableCell align="center">ลบข้อมูล</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>

         {carcheckinout.map((item:any )=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.edges.Ambulance.carregistration}</TableCell>
             <TableCell align="center">{item.edges.Name.name}</TableCell>
             <TableCell align="center">{item.edges.Purpose.objective}</TableCell>
             <TableCell align="center">{item.note}</TableCell>
             <TableCell align="center">{item.person}</TableCell>
             <TableCell align="center">{item.distance}</TableCell>
             <TableCell align="center">{item.place}</TableCell>
             <TableCell align="center">{moment(item.checkOut).format('DD/MM/YYYY HH.mm น.')}</TableCell>
             <TableCell align="center">{moment(item.checkIn).format('DD/MM/YYYY HH.mm น.')}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteCarcheckinouts(item.id);
                 }}
                 style={{ marginLeft: 2 }}
                 variant="contained" 
                 color="secondary"
                 startIcon={<DeleteForeverIcon/>}
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Content>
 );

}
