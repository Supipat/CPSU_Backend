// src/pages/Document.js
import React, { useState } from "react";
import Navbar from "../component/navbar";
import Headers from "../component/header";
import Footer from "../component/footer";
import "../css/document.css";

const petitionGroups = {
  "แบบฟอร์มทั่วไป": [
    { name: "คำร้องขอตรวจสอบคะแนนสอบ", link: "https://www.sc.su.ac.th/FormPDF/score_form.pdf" },
    { name: "คำร้องขอใช้สถานที่", link: "https://www.sc.su.ac.th/FormPDF/place_form.pdf" },
    { name: "คำร้องขอผ่อนผันการชำระเงินค่าลงทะเบียน", link: "https://www.sc.su.ac.th/FormPDF/register.pdf" },
    { name: "คำร้องขอเปลี่ยนสาขาวิชาเอก", link: "https://www.sc.su.ac.th/FormPDF/major_form.pdf" },
    { name: "คำร้องขอขาดเรียนและขาดสอบ", link: "https://www.sc.su.ac.th/FormPDF/double.pdf" },
    { name: "คำร้องขอสำรองที่นั่ง", link: "https://www.sc.su.ac.th/FormPDF/absent.pdf" },
    { name: "คำร้องขอกลับเข้าศึกษา", link: "https://www.sc.su.ac.th/FormPDF/booking.pdf" },
    { name: "คำร้องทั่วไป", link: "https://www.sc.su.ac.th/FormPDF/generalForm.pdf" },
    { name: "คำร้องขอลงทะเบียน/เพิ่ม/ถอน ช้าเป็นกรณีพิเศษ", link: "https://edoc.sc.su.ac.th/StuDoc.php?docid=9" },
   ],
  "แบบฟอร์มโครงงานปริญญานิพนธ์(1)": [
    { name: "CP00R แบบคำร้องขอเข้ารับการจัดสรรกลุ่มสำหรับการจัดทำโครงงานปริญญานิพนธ์", link: "https://cp.su.ac.th/file/show/54" },
    { name: "CP00X แบบแจ้งความประสงค์แลกเปลี่ยนกลุ่มโครงงานปริญญานิพนธ์", link: "https://cp.su.ac.th/file/show/55" },
    { name: "CS00G แบบขอตรวจสอบคุณสมบัติในการมีสิทธิขอจัดทําโครงงานปริญญานิพนธ์ (สาขาวิชาวิทยาการคอมพิวเตอร์) (หลักสูตร 51)", link: "https://cp.su.ac.th/file/show/56" },
    { name: "CS00G-55 แบบขอตรวจสอบคุณสมบัติในการมีสิทธิขอจัดทําโครงงานปริญญานิพนธ์ (สาขาวิชาวิทยาการคอมพิวเตอร์) (หลักสูตร 55)", link: "https://cp.su.ac.th/file/show/57" },
    { name: "IT00G แบบขอตรวจสอบคุณสมบัติในการมีสิทธิขอจัดทําโครงงานปริญญานิพนธ์ (สาขาวิชาเทคโนโลยีสารสนเทศ) (หลักสูตร 51)", link: "https://cp.su.ac.th/file/show/58" },
    { name: "IT00G-55 แบบขอตรวจสอบคุณสมบัติในการมีสิทธิขอจัดทําโครงงานปริญญานิพนธ์ (สาขาวิชาเทคโนโลยีสารสนเทศ) (หลักสูตร 55)", link: "https://cp.su.ac.th/file/show/59" },
    { name: "CS01S/IT01S แบบคําร้องขอเสนอหัวข้อโครงงานปริญญานิพนธ์", link: "https://cp.su.ac.th/file/show/60" },
    { name: "CS01D/IT01D Template ไฟล์ต้นแบบเอกสารประกอบการเสนอหัวข้อโครงงานปริญญานิพนธ์", link: "https://cp.su.ac.th/file/show/61" },
    { name: "CS01D/IT01D Example ตัวอย่างเอกสารประกอบการเสนอหัวข้อโครงงานปริญญานิพนธ์", link: "https://cp.su.ac.th/file/show/62" },
    { name: "CS02S/IT02S แบบคําร้องขอสอบข้อเสนอโครงงานปริญญานิพนธ์", link: "https://cp.su.ac.th/file/show/63" },
    { name: "CS02D/IT02D Template ไฟล์ต้นแบบเอกสารประกอบการสอบข้อเสนอโครงงานปริญญานิพนธ์", link: "https://cp.su.ac.th/file/show/64" },
    { name: "CS02R/IT02R แบบประเมินผลการสอบข้อเสนอโครงงานปริญญานิพพนธ์", link: "https://cp.su.ac.th/file/show/65" },
  ],
  "แบบฟอร์มโครงงานปริญญานิพนธ์(2)": [
    { name: "CS03S/IT03S แบบคําร้องขอสอบติดตามความก้าวหน้าโครงงานปริญญานิพนธ์", link: "https://cp.su.ac.th/file/show/67" },
    { name: "CS03D/IT03D Template ไฟล์ต้นแบบเอกสารประกอบการสอบติดตามความก้าวหน้าโครงงานปริญญานิพนธ์", link: "https://cp.su.ac.th/file/show/68" },
    { name: "CS03R/IT03R แบบประเมินผลการสอบติดตามความก้าวหน้าโครงงานปริญญานิพนธ์", link: "https://cp.su.ac.th/file/show/69" },
    { name: "CS04S/IT04S แบบคำร้องขอสอบนำเสนอโครงงานปริญญานิพนธ์ที่เสร็จสมบูรณ์", link: "https://cp.su.ac.th/file/show/70" },
    { name: "CS04D/IT04D Template ไฟล์ต้นแบบเอกสารประกอบการสอบโครงงานปริญญานิพนธ์ที่เสร็จสมบูรณ์", link: "https://cp.su.ac.th/file/show/71" },
    { name: "CS04R/IT04R แบบประเมินผลการสอบนำเสนอโครงงานปริญญานิพนธ์ที่เสร็จสมบูรณ์", link: "https://cp.su.ac.th/file/show/72" },
    { name: "CS05D/IT05D Template ไฟล์ต้นแบบเอกสารโครงงานปริญญานิพนธ์ที่เสร็จสมบูรณ์", link: "https://cp.su.ac.th/file/show/73" },
    { name: "CS05D/IT05D CD Template ไฟล์ต้นแบบแผ่นหน้าปกซีดีโครงงานปริญญานิพนธ์ที่เสร็จสมบูรณ์", link: "https://cp.su.ac.th/file/show/74" },
    { name: "CS05C/IT05C แบบฟอร์มการส่งตรวจรูปแบบเล่มโครงงานปริญญานิพนธ์", link: "https://cp.su.ac.th/file/show/75" },
    { name: "ไฟล์ต้นแบบโปสเตอร์ (Poster Template)", link: "https://cp.su.ac.th/file/show/338" },
  ],
};

const Document = () => {
  const [searchTerm, setSearchTerm] = useState("");

  const handleSearch = (e) => {
    setSearchTerm(e.target.value.toLowerCase());
  };

  const filteredGroups = Object.entries(petitionGroups)
    .map(([groupName, petitions]) => {
      const filteredPetitions = petitions.filter((petition) =>
        petition.name.toLowerCase().includes(searchTerm)
      );
      return { groupName, petitions: filteredPetitions };
    })
    .filter((group) => group.petitions.length > 0);

  return (
    <>
      <Headers />
      <Navbar />

      <div className="container document-container my-5">
        <h2 className="section-title text-center mb-4">เอกสารคำร้องและแบบฟอร์ม</h2>

        <input
          type="text"
          className="form-control mb-4"
          placeholder="ค้นหาเอกสาร..."
          onChange={handleSearch}
          value={searchTerm}
        />

        {filteredGroups.map(({ groupName, petitions }) => (
          <section key={groupName} className="mb-5">
            <h3 className="text-primary mb-3">{groupName}</h3>
            <ul className="list-group">
              {petitions.map((petition, index) => (
                <li key={index} className="list-group-item">
                  <a href={petition.link} target="_blank" rel="noopener noreferrer">
                    {petition.name}
                  </a>
                </li>
              ))}
            </ul>
          </section>
        ))}

        {filteredGroups.length === 0 && (
          <p className="text-center text-muted">ไม่พบเอกสารที่ตรงกับคำค้นหา</p>
        )}
      </div>

      <Footer />
    </>
  );
};

export default Document;
