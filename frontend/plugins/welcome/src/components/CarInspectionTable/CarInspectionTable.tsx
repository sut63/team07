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

import { EntCarInspection } from '../../api/models/EntCarInspection';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [carinspections, setCarInspections] = useState<EntCarInspection[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getCarInspections = async () => {
     const res = await api.listCarinspection();
     setLoading(false);
     setCarInspections(res);
   };
   getCarInspections();

 }, [loading]);
 
 const deleteCarInspections = async (id: number) => {
   const res = await api.deleteCarinspection({ id: id });
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
           <TableCell align="center">ผลการตรวจสภาพรถ</TableCell>
           <TableCell align="center">วัน/เดือน/ปี เวลา</TableCell>
           <TableCell align="center">หมายเหตุ</TableCell>
           <TableCell align="center">จัดการข้อมูล</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {carinspections.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges.user.name}</TableCell>
             <TableCell align="center">{item.edges.ambulance.carregistration}</TableCell>
             <TableCell align="center">{item.edges.inspectionresult.resultName}</TableCell>
             <TableCell align="center">{moment(item.datetime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
             <TableCell align="center">{item.note}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                    deleteCarInspections(item.id);
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