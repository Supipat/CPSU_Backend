-- create news

CREATE TABLE IF NOT EXISTS news_types (
    type_id SERIAL PRIMARY KEY,
    type_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS news (
    news_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    type_id INT NOT NULL,
    detail_url TEXT NOT NULL,
    cover_image TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (type_id) REFERENCES news_types(type_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS news_images (
    image_id SERIAL PRIMARY KEY,
    news_id INT NOT NULL,
    file_image TEXT NOT NULL,
    FOREIGN KEY (news_id) REFERENCES news(news_id) ON DELETE CASCADE
);

-- create courses

CREATE TABLE IF NOT EXISTS career_paths (
    career_paths_id SERIAL PRIMARY KEY,
    career_paths TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS plo (
    plo_id SERIAL PRIMARY KEY,
    plo TEXT NOT NULL
);

-- ระดับปริญญา + สาขา + ปี

CREATE TABLE IF NOT EXISTS courses (
    course_id VARCHAR(10) PRIMARY KEY,
    degree VARCHAR(100) NOT NULL,
    major VARCHAR(100) NOT NULL,
    year INT NOT NULL,
    thai_course VARCHAR(255) NOT NULL,
    eng_course VARCHAR(255) NOT NULL,
    thai_degree VARCHAR(255) NOT NULL,
    eng_degree VARCHAR(255) NOT NULL,
    admission_req TEXT NOT NULL,
    graduation_req TEXT NOT NULL,
    philosophy TEXT NOT NULL,
    objective TEXT NOT NULL,
    tuition TEXT NOT NULL,
    credits VARCHAR(50) NOT NULL,
    career_paths_id INT NOT NULL,
    plo_id INT NOT NULL,
    detail_url TEXT NOT NULL,
    status VARCHAR(25) NOT NULL,
    FOREIGN KEY (career_paths_id) REFERENCES career_paths(career_paths_id) ON DELETE CASCADE,
    FOREIGN KEY (plo_id) REFERENCES plo(plo_id) ON DELETE CASCADE
);

-- create course structure

CREATE TABLE IF NOT EXISTS course_structure (
    course_structure_id SERIAL PRIMARY KEY,
    course_id VARCHAR(10) NOT NULL,
    course_structure_url TEXT NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses(course_id) ON DELETE CASCADE
);

-- create roadmap

CREATE TABLE IF NOT EXISTS roadmap (
    roadmap_id SERIAL PRIMARY KEY,
    course_id VARCHAR(10) NOT NULL,
    roadmap_url TEXT NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses(course_id) ON DELETE CASCADE
);

-- create subject

CREATE TABLE IF NOT EXISTS description (
    description_id VARCHAR(6) PRIMARY KEY,
    description_thai TEXT NULL,
    description_eng TEXT NULL
);

CREATE TABLE IF NOT EXISTS clo (
    clo_id VARCHAR(6) PRIMARY KEY,
    clo TEXT NULL
);

CREATE TABLE IF NOT EXISTS subjects (
    id SERIAL PRIMARY KEY,
    subject_id VARCHAR(10) NOT NULL,
    course_id VARCHAR(10) NOT NULL,
    plan_type VARCHAR(50) NOT NULL,
    semester VARCHAR(50) NOT NULL,
    thai_subject VARCHAR(100) NOT NULL,
    eng_subject VARCHAR(100) NULL,
    credits VARCHAR(50) NOT NULL,
    compulsory_subject VARCHAR(255) NULL,
    condition VARCHAR(255) NULL,
    description_id VARCHAR(6) NULL,
    clo_id VARCHAR(6) NULL,
    FOREIGN KEY (course_id) REFERENCES courses(course_id) ON DELETE CASCADE,
    FOREIGN KEY (description_id) REFERENCES description(description_id) ON DELETE CASCADE,
    FOREIGN KEY (clo_id) REFERENCES clo(clo_id) ON DELETE CASCADE
);



-- create personnel

CREATE TABLE IF NOT EXISTS department_position (
    department_position_id SERIAL PRIMARY KEY,
    department_position_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS academic_position (
    academic_position_id SERIAL PRIMARY KEY,
    thai_academic_position VARCHAR(50) NOT NULL,
    eng_academic_position VARCHAR(50) NULL
);

CREATE TABLE IF NOT EXISTS personnels (
    personnel_id SERIAL PRIMARY KEY,
    type_personnel VARCHAR(50) NOT NULL,
    department_position_id INT NOT NULL,
    academic_position_id INT NULL,
    thai_name VARCHAR(50) NOT NULL,
    eng_name VARCHAR(50) NOT NULL,
    education TEXT NULL,
    related_fields TEXT NULL,
    email VARCHAR(100) NULL,
    website TEXT NULL,
    file_image TEXT NOT NULL,
    scopus_id VARCHAR(50) NULL,
    FOREIGN KEY (department_position_id) REFERENCES department_position(department_position_id) ON DELETE CASCADE,
    FOREIGN KEY (academic_position_id) REFERENCES academic_position(academic_position_id) ON DELETE CASCADE
);

-- create research

CREATE TABLE IF NOT EXISTS research (
    research_id SERIAL PRIMARY KEY,
    personnel_id INT NOT NULL,
    title TEXT NOT NULL,                                       
    journal VARCHAR(255) NOT NULL,              
    year INT NOT NULL,                         
    volume VARCHAR(50) NULL,                         
    issue VARCHAR(50) NULL,                         
    pages VARCHAR(50) NULL,                          
    doi TEXT NULL,                          
    cited INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (personnel_id) REFERENCES personnels(personnel_id) ON DELETE CASCADE           
);

-- FOREIGN KEY (type_id) REFERENCES type_personnel(type_id) ON DELETE CASCADE,

-- TRIGGER

CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now(); 
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_news_modtime
BEFORE UPDATE ON news
FOR EACH ROW
EXECUTE PROCEDURE update_modified_column();

-- insert news

INSERT INTO news_types(type_name) VALUES
('ข่าวประชาสัมพันธ์'),
('ทุนการศึกษา'),
('รางวัลที่ได้รับ'),
('กิจกรรมของภาควิชา');

INSERT INTO news(title,content,type_id,detail_url,cover_image) VALUES
('คู่มือแนะนำนักศึกษาใหม่ ปีการศึกษา 2568',
    'คู่มือแนะนำ การลงทะเบียนรายวิชาเรียน การเพิ่ม - ถอน - การเปลี่ยนกลุ่มเรียน ในรายวิชาเดิม การถอนรายวิชาโดยติดสัญลักษณ์ W ช่องทางการชำระค่าธรรมเนียมการศึกษา เอกสารแนบเบิกค่าเล่าเรียน การลาพักการศึกษา นักศึกษาที่คาดว่าจะสำเร็จการศึกษา การขอหนังสือสำคัญทางการศึกษา การติดต่อเจ้าหน้าที่งานทะเบียนและประมวลผล ช่องทางติดต่อคณะวิชา กองกิจการนักศึกษา การให้บริการสำหรับนักศึกษา วิทยาเขตสารสนเทศเพชรบุรี สำนักดิจิทัลเทคโนโลยี หอสมุด มหาวิทยาลัยศิลปากร ประกันอุบัติเหตุส่วนบุคคล ศูนย์บริการสุขภาพ คณะเภสัชศาสตร์ บริการด้านสุขภาพร่างกาย เครื่องแบบนักศึกษา / บัตรนักศึกษาอิเล็กทรอนิกส์',
    (SELECT type_id FROM news_types WHERE type_name = 'ข่าวประชาสัมพันธ์' LIMIT 1),
    'https://drive.google.com/file/d/1FnJRketlluku27HQqs-mVPnxMJ5wFxqZ/view?usp=sharing&fbclid=IwZXh0bgNhZW0CMTAAAR3M6TbS4tN1DUSa-z2NaM1ekAOELZp7TnsIhJWC6g_dvfz-sD_b0La0S7U_aem_K5_Zhi2eLKhIpJxaeszOlQ',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/manual.jpg'),
('เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล',
    'เปิดรับสมัครแล้ว! ภาควิชาคอมพิวเตอร์ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร เปิดรับสมัครนักศึกษาระดับบัณฑิตศึกษา ปริญญาโท - ปริญญาเอก หลักสูตร IT สำหรับคนทำงาน สู่ผู้เชี่ยวชาญด้านวิจัยและนวัตกรรม เน้นเรียนรู้กระบวนทำวิจัยเพื่อแก้ปัญหาจริง มีทุนผู้ช่วยวิจัยและทุนนำเสนอผลงานในงานประชุมวิชาการ มีทั้งภาคปกติ (เรียนวันธรรมดา) และโครงการพิเศษ (เรียน ส อา) มี 3 แผนการเรียนให้เลือก Data Science, Project Management, DevOps เปิดรับสมัครรอบ 1 วันที่ 16 ธ.ค. 67 - 21 ก.พ. 68 เปิดรับสมัครรอบ 2 วันี่ต่ 3 มี.ค. 68 - 7 พ.ค. 68 สมัครง่าย ๆ ผ่านช่องทางออนไลน์ได้ที่ https://graduate.su.ac.th มีข้อสงสัยสอบถามเพิ่มเติมได้ที่ เพจ Facebook : https://www.facebook.com/computingsu/ หรือ โทร 034-272-923',
    (SELECT type_id FROM news_types WHERE type_name = 'ข่าวประชาสัมพันธ์' LIMIT 1),
    'https://graduate.su.ac.th',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/degree1.jpg'),
('กำหนดการรับสมัครและสัมภาษณ์ทุนการศึกษา ประจำปีการศึกษา 2568 (ครั้งที่ 1)',
    'กำหนดการรับสมัครและสัมภาษณ์ทุนการศึกษา ของคณะวิทยาศาสตร์ เพื่อขอทุนการศึกษาสำหรับนักศึกษาปริญญาตรี ภาคการศึกษาต้น ปีการศึกษา 2568 สังกัดคณะวิทย์ ขาดแคลนทุน เกรดไม่ต่ำกว่า 2.00 ขอและยื่นใบสัมครได้ที่ " งานบริการการศึกษา ชั้น 1 อาคารวิทย์ 1 "',
    (SELECT type_id FROM news_types WHERE type_name = 'ทุนการศึกษา' LIMIT 1),
    'https://dsa.su.ac.th/ksu/wp-content/uploads/2025/06/%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B8%81%E0%B8%B2%E0%B8%A8%E0%B8%81%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B4%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%99%E0%B8%B1%E0%B8%81%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2-%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87-%E0%B8%81%E0%B8%B3%E0%B8%AB%E0%B8%99%E0%B8%94%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A3.pdf',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/scholarship.jpg'),
('ขอแสดงความยินดีกับ ผู้ช่วยศาสตราจารย์ ดร.อรวรรณ เชาวลิต อาจารย์ประจำภาควิชาคอมพิวเตอร์',
    'คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร ขอแสดงความยินดีกับ ผู้ช่วยศาสตราจารย์ ดร.อรวรรณ เชาวลิต ภาควิชาคอมพิวเตอร์ ในโอกาสที่ตีพิมพ์ผลงานวิจัยในวารสาร ICIC Express Letters, Part B: Applications ในฐานข้อมูล Scopus (ScimagoJR, Quartile 4) เรื่อง Multivariate time series forecasting Thailand soybean meal price with deep learning models ICIC Express Letters, Part B: Applications 2025, 16(3), 317-323. http://www.icicelb.org/ellb/contents/2025/3/elb-16-03-10.pdf SDG 8 DECENT WORK AND ECONOMIC GROWTH',
    (SELECT type_id FROM news_types WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'http://www.icicelb.org/ellb/contents/2025/3/elb-16-03-10.pdf',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-orawan.jpg'),
('ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์ ดร.ปัญญนัท อ้นพงษ์ เนื่องในโอกาสได้รับการแต่งตั้งให้ดำรงตำแหน่งทางวิชาการ',
    'ภาควิชาคอมพิวเตอร์ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์ ดร.ปัญญนัท อ้นพงษ์ เนื่องในโอกาสได้รับการแต่งตั้งให้ดำรงตำแหน่งทางวิชาการ ตำแหน่ง “ผู้ช่วยศาสตราจารย์” ในสาขาวิชาเทคโนโลยีสารสนเทศ (1806)',
    (SELECT type_id FROM news_types WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'https://www.facebook.com/ScienceSilpakornUniversity/posts/คณะวิทยาศาสตร์-มหาวิทยาลัยศิลปากร-ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์-ดรปัญญนัท/982826210555894/',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-panyanut.jpg'),
('ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล',
    'ภาควิชาคอมพิวเตอร์ ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล ที่ได้รับรางวัล Best Presentation ในงาน The 6th Asia Joint Conference on Computing (AJCC 2025) ณ เมือง Osaka ประเทศญี่ปุ่นจาก Paper หัวข้อ "Intention Classification of Chinese Topics on Thai Facebook Pages Using Transformer Models with Emotional Features" อาจารย์ที่ปรึกษา อ.ดร.สัจจาภรณ์ ไวจรรยา และ ผศ.ดร.ณัฐโชติ พรหมฤทธิ์',
    (SELECT type_id FROM news_types WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'https://www.facebook.com/computingsu/posts/-ภาควิชาคอมพิวเตอร์-ขอแสดงความยินดีกับ-นายภากร-กัทชลี-นักศึกษา-ปริญญาเอก-หลักสูต/122207080106177331/',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-phakon1.jpg'),
('ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร',
    'ภาควิชาคอมพิวเตอร์ ขอแสดงความยินดีกับนักศึกษาและอาจารย์ที่ปรึกษา ที่เข้าร่วมนำเสนอผลงานและได้รับรางวัลในงาน การประชุมวิชาการระดับปริญญาตรีด้านคอมพิวเตอร์ภูมิภาคเอเชีย ครั้งที่ 13 (The 13th Asia Undergraduate Conference on Computing: AUCC 2025) ในระหว่างวันที่ 19 - 21 กุมภาพันธ์ พ.ศ. 2568 ณ มหาวิทยาลัยราชภัฏอุตรดิตถ์ จังหวัดอุตรดิตถ์ รางวัลรองชนะเลิศอันดับ 1 Track นวัตกรรม กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร (Development of examination Management system for the Faculty of Science)" ซึ่งผลงานนี้ได้มีการนำไปใช้จริงทดแทนระบบเดิม ในการจัดสอบของคณะฯ เมื่อช่วงสอบกลางภาคการศึกษาที่ผ่านมา นักศึกษาที่เข้านำเสนอ - นาย ณัฐพิสิษ กังวานธรรมกุล - นาย นัทธพงศ์ เป็กทอง - นาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ อาจารย์ที่ปรึกษา ผศ.ดร.ณัฐโชติ พรหมฤทธิ์ และ อ.ดร.สัจจาภรณ์ ไวจรรยา',
    (SELECT type_id FROM news_types WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'https://www.facebook.com/computingsu/posts/-ภาควิชาคอมพิวเตอร์-ขอแสดงความยินดีกับนักศึกษาและอาจารย์ที่ปรึกษา-ที่เข้าร่วมนำเ/122194936622177331/',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward2nd1.jpg'),
('งานสานสัมพันธ์ภาคคอมพิวเตอร์',
    'เตรียมตัวให้พร้อม! งานสานสัมพันธ์ภาคคอมพิวเตอร์กำลังจะเริ่มขึ้น! เจอกันวันที่ 28 กุมภาพันธ์ 2568 ณ ลานจอดรถตึก 4 เวลา 17:30 - 21:30 น. ธีม: ย้อนยุค ไม่จำกัดช่วงเวลา! จะเป็นยุคหิน มนุษย์ถ้ำ อารยธรรมโบราณ หรือยุคใดก็ได้ จัดเต็มมาให้สุด! มีรางวัลแต่งกายยอดเยี่ยม 3 รางวัล!ใครแต่งตัวเข้าธีมและโดดเด่นที่สุด มีสิทธิ์คว้ารางวัลไปเลย! สนุกไปกับการแสดงสุดพิเศษ, เกมสุดมันส์, ลุ้นรางวัล และดนตรีสดปิดท้าย! ใครลงชื่อไว้แล้ว เจอกันแน่นอน! เตรียมชุดให้พร้อม แล้วมาสนุกไปด้วยกัน!',
    (SELECT type_id FROM news_types WHERE type_name = 'กิจกรรมของภาควิชา' LIMIT 1),
    '',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/cpsu_event.png');

INSERT INTO news_images(news_id,file_image) VALUES
((SELECT news_id FROM news WHERE title = 'คู่มือแนะนำนักศึกษาใหม่ ปีการศึกษา 2568' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/manual.jpg'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/degree1.jpg'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/degree2.jpg'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/degree3.jpg'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/degree4.jpg'),
((SELECT news_id FROM news WHERE title = 'กำหนดการรับสมัครและสัมภาษณ์ทุนการศึกษา ประจำปีการศึกษา 2568 (ครั้งที่ 1)' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/scholarship.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ ผู้ช่วยศาสตราจารย์ ดร.อรวรรณ เชาวลิต อาจารย์ประจำภาควิชาคอมพิวเตอร์' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-orawan.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์ ดร.ปัญญนัท อ้นพงษ์ เนื่องในโอกาสได้รับการแต่งตั้งให้ดำรงตำแหน่งทางวิชาการ' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-panyanut.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-phakon1.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-phakon2.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward2nd1.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward2nd2.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward2nd3.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward2nd4.jpg'),
((SELECT news_id FROM news WHERE title = 'งานสานสัมพันธ์ภาคคอมพิวเตอร์' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/cpsu_event.png');

-- insert courses

COPY career_paths(career_paths)
FROM '/docker-entrypoint-initdb.d/csv/course/career_paths.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

COPY plo(plo)
FROM '/docker-entrypoint-initdb.d/csv/course/plo.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

COPY courses(course_id,degree,major,year,thai_course,eng_course,thai_degree,eng_degree,admission_req,graduation_req,philosophy,objective,tuition,credits,career_paths_id,plo_id,detail_url,status)
FROM '/docker-entrypoint-initdb.d/csv/course/courses.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

SELECT setval('career_paths_career_paths_id_seq', (SELECT MAX(career_paths_id) FROM career_paths));
SELECT setval('plo_plo_id_seq', (SELECT MAX(plo_id) FROM plo));

-- insert course structure

INSERT INTO course_structure(course_id,course_structure_url) VALUES
((SELECT course_id FROM courses WHERE thai_course = '(วท.บ) หลักสูตรวิทยาศาสตรบัณฑิต สาขาวิชาวิทยาการคอมพิวเตอร์ 2565' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/course/course_structure_CS_65.png'),
((SELECT course_id FROM courses WHERE thai_course = '(วท.บ) หลักสูตรวิทยาศาสตรบัณฑิต สาขาวิชาเทคโนโลยีสารสนเทศ 2565' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/course/course_structure_IT_65.png');

-- insert roadmap

INSERT INTO roadmap(course_id,roadmap_url) VALUES
((SELECT course_id FROM courses WHERE thai_course = '(วท.บ) หลักสูตรวิทยาศาสตรบัณฑิต สาขาวิชาวิทยาการคอมพิวเตอร์ 2565' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/course/roadmap_CS_65.jpg'),
((SELECT course_id FROM courses WHERE thai_course = '(วท.บ) หลักสูตรวิทยาศาสตรบัณฑิต สาขาวิชาเทคโนโลยีสารสนเทศ 2565' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/course/roadmap_IT_65.jpg');

-- insert subject

COPY description(description_id,description_thai,description_eng)
FROM '/docker-entrypoint-initdb.d/csv/subject/BS_description.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

COPY description(description_id,description_thai,description_eng)
FROM '/docker-entrypoint-initdb.d/csv/subject/MSIT66_description.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

COPY description(description_id,description_thai,description_eng)
FROM '/docker-entrypoint-initdb.d/csv/subject/DSIT66_description.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

COPY clo(clo_id,clo)
FROM '/docker-entrypoint-initdb.d/csv/subject/BS_clo.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

COPY clo(clo_id,clo)
FROM '/docker-entrypoint-initdb.d/csv/subject/MSIT66_clo.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

COPY clo(clo_id,clo)
FROM '/docker-entrypoint-initdb.d/csv/subject/DSIT66_clo.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

COPY subjects(subject_id,course_id,plan_type,semester,thai_subject,eng_subject,credits,compulsory_subject,condition,description_id,clo_id)
FROM '/docker-entrypoint-initdb.d/csv/subject/BS_subjects.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

COPY subjects(subject_id,course_id,plan_type,semester,thai_subject,eng_subject,credits,compulsory_subject,condition,description_id,clo_id)
FROM '/docker-entrypoint-initdb.d/csv/subject/MSIT66_subjects.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

COPY subjects(subject_id,course_id,plan_type,semester,thai_subject,eng_subject,credits,compulsory_subject,condition,description_id,clo_id)
FROM '/docker-entrypoint-initdb.d/csv/subject/DSIT66_subjects.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

SELECT setval('subjects_id_seq', (SELECT MAX(id) FROM subjects));

-- insert personnel

INSERT INTO department_position(department_position_name) VALUES
('หัวหน้าภาควิชา'),('รองหัวหน้าภาควิชาฯ ฝ่ายบริหาร'),('รองหัวหน้าภาควิชาฯ'),('อาจารย์ประจำภาควิชา'),
('นักวิชาการอุดมศึกษาชำนาญการ'),('นักวิชาการอุดมศึกษาปฏิบัติการ'),('นักวิชาการอุดมศึกษา (ประจำหลักสูตรวิทยาการข้อมูล)'),('นักวิชาการอุดมศึกษา'),('นักเทคโนโลยีสารสนเทศ'),('นักคอมพิวเตอร์'),('พนักงานทั่วไป');

INSERT INTO academic_position(thai_academic_position,eng_academic_position) VALUES
('รศ.ดร.','Assoc.Prof.Dr.'),
('ผศ.ดร.','Asst.Prof.Dr.'),
('ผศ.','Asst.Prof.'),
('อ.ดร.','Dr.'),
('อ.','');

COPY personnels(type_personnel,department_position_id,academic_position_id,thai_name,eng_name,education,related_fields,email,website,file_image,scopus_id)
FROM '/docker-entrypoint-initdb.d/csv/personnel/personnels.csv'
WITH (FORMAT csv, HEADER true, ENCODING 'UTF8');

SELECT setval('personnels_personnel_id_seq', (SELECT MAX(personnel_id) FROM personnels));