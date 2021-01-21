import React, { useState, useEffect } from 'react';
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
import { EntAmbulance } from '../../api/models/EntAmbulance';

import moment from 'moment';

const useStyles = makeStyles({
 table: {
   minWidth: 1500,
 },
});
export default function ComponentsTable() {
 
 const classes = useStyles();
 const api = new DefaultApi();
 const [loading, setLoading] = useState(true);
 const [ambulance, setAmbulance] = useState<EntAmbulance[]>([]);
 const [userid, setUser] = useState(Number);

 useEffect(() => {
   const getAmbulances = async () => {
     const res = await api.listAmbulance({ limit: 120, offset: 0 });
     setLoading(false);
     setAmbulance(res);
   };
   getAmbulances();
   const checkJobPosition = async () => {
    const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
    setLoading(false);
    if (jobdata != "เจ้าหน้าที่รถพยาบาล" ) {
      localStorage.setItem("userdata",JSON.stringify(null));
      localStorage.setItem("jobpositiondata",JSON.stringify(null));
      history.pushState("","","./");
      window.location.reload(false);        
    }
    else{
        setUser(Number(localStorage.getItem("userdata")))
    }
  }
checkJobPosition();
 }, [loading]);

 const deleteAmbulances = async (id: number) => {
    await api.deleteAmbulance({ id: id });
   setLoading(true);
 };

 return (
   

   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">ลำดับ</TableCell>
           <TableCell align="center">ทะเบียนรถ</TableCell>
           <TableCell align="center">แบรนด์รถ</TableCell>
           <TableCell align="center">เครื่องยนต์</TableCell>
           <TableCell align="center">ความจุตัวถัง</TableCell>
           <TableCell align="center">บริษัทประกัน</TableCell>
           <TableCell align="center">สถานะ</TableCell>
           <TableCell align="center">เจ้าหน้าที่</TableCell>
           <TableCell align="center">วันที่</TableCell>
           <TableCell align="center">ลบข้อมูล</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>

         {ambulance.map((item:any )=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.carregistration}</TableCell>
             <TableCell align="center">{item.edges?.Hasbrand?.brand}</TableCell>
             <TableCell align="center">{item.enginepower}</TableCell>
             <TableCell align="center">{item.displacement}</TableCell>
             <TableCell align="center">{item.edges?.Hasinsurance?.company}</TableCell>
             <TableCell align="center">{item.edges?.Hasstatus?.resultName}</TableCell>
             <TableCell align="center">{item.edges?.Hasuser?.name}</TableCell>
             <TableCell align="center">{moment(item.registerat).format('DD/MM/YYYY HH.mm น.')}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteAmbulances(item.id);
                 }}
                 style={{ marginLeft: 2 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>

 );

}
