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
import { EntCarCheckInOut } from '../../api/models/EntCarCheckInOut';
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
 const [carcheckinout, setCarcheckinout] = useState<EntCarCheckInOut[]>([]);

 useEffect(() => {
   const getCarcheckinouts = async () => {
     const res = await api.listCarcheckinout({ limit: 120, offset: 0 });
     setLoading(false);
     setCarcheckinout(res);
   };
   getCarcheckinouts();
 }, [loading]);

 const deleteCarcheckinouts = async (id: number) => {
    await api.deleteCarcheckinout({ id: id });
   setLoading(true);
 };

 return (
   

   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">ลำดับ</TableCell>
           <TableCell align="center">ทะเบียนรถ</TableCell>
           <TableCell align="center">เจ้าหน้าที่</TableCell>
           <TableCell align="center">วัตถุประสงค์</TableCell>
           <TableCell align="center">หมายเหตุ</TableCell>
           <TableCell align="center">วันที่รถออก</TableCell>
           <TableCell align="center">วันที่รถเข้า</TableCell>
           <TableCell align="center">ลบข้อมูล</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>

         {carcheckinout.map((item:any )=> (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges.ambulance.carregistration}</TableCell>
             <TableCell align="center">{item.edges.name.name}</TableCell>
             <TableCell align="center">{item.edges.purpose.objective}</TableCell>
             <TableCell align="center">{item.note}</TableCell>
             <TableCell align="center">{moment(item.checkOut).format('DD/MM/YYYY HH.mm น.')}</TableCell>
             <TableCell align="center">{moment(item.checkIn).format('DD/MM/YYYY HH.mm น.')}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   deleteCarcheckinouts(item.id);
                 }}
                 style={{ marginLeft: 2 }}
                 variant="contained" 
                 color="primary"
                 style={{backgroundColor: "#F5454D"}}
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
