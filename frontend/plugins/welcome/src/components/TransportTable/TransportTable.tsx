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

import moment from 'moment';

import { EntTransport } from '../../api/models/EntTransport';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [transports, setTransports] = useState<EntTransport[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getTransports = async () => {
     const res = await api.listTransport();
     setLoading(false);
     setTransports(res);
   };
   getTransports();

 }, [loading]);
 
 const deleteTransports = async (id: number) => {
   const res = await api.deleteTransport({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="center">เลขที่</TableCell>
           <TableCell align="center">เจ้าหน้าที่</TableCell>
           <TableCell align="center">รถพยาบาล</TableCell>
           <TableCell align="center">โรงพยาบาลที่จะส่ง</TableCell>
           <TableCell align="center">โรงพยาบาลที่จะรับ</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {transports.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges.user.name}</TableCell>
             <TableCell align="center">{item.edges.ambulance.carregistration}</TableCell>
             <TableCell align="center">{item.edges.send.sendname}</TableCell>
             <TableCell align="center">{item.edges.receive.receivename}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                    deleteTransports(item.id);
                 }}
                 style={{ marginLeft: 10 , backgroundColor: "#140087" }}
                 variant="contained"
                 color="secondary"
               >
                 ลบข้อมูล
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}