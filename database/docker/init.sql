-- CREATE news

CREATE TABLE IF NOT EXISTS news_type (
    type_id SERIAL PRIMARY KEY,
    type_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS news (
    news_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    type_id INT NOT NULL,
    detail_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (type_id) REFERENCES news_type(type_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS news_image (
    image_id SERIAL PRIMARY KEY,
    news_id INT NOT NULL,
    file_image TEXT NOT NULL,
    FOREIGN KEY (news_id) REFERENCES news(news_id) ON DELETE CASCADE
);

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

-- INSERT news

INSERT INTO news_type(type_name) VALUES
('ข่าวประชาสัมพันธ์'),
('ทุนการศึกษา'),
('รางวัลที่ได้รับ'),
('กิจกรรมของภาควิชา');

INSERT INTO news(title,content,type_id,detail_url) VALUES
('คู่มือแนะนำนักศึกษาใหม่ ปีการศึกษา 2568',
    'คู่มือแนะนำ การลงทะเบียนรายวิชาเรียน การเพิ่ม - ถอน - การเปลี่ยนกลุ่มเรียน ในรายวิชาเดิม การถอนรายวิชาโดยติดสัญลักษณ์ W ช่องทางการชำระค่าธรรมเนียมการศึกษา เอกสารแนบเบิกค่าเล่าเรียน การลาพักการศึกษา นักศึกษาที่คาดว่าจะสำเร็จการศึกษา การขอหนังสือสำคัญทางการศึกษา การติดต่อเจ้าหน้าที่งานทะเบียนและประมวลผล ช่องทางติดต่อคณะวิชา กองกิจการนักศึกษา การให้บริการสำหรับนักศึกษา วิทยาเขตสารสนเทศเพชรบุรี สำนักดิจิทัลเทคโนโลยี หอสมุด มหาวิทยาลัยศิลปากร ประกันอุบัติเหตุส่วนบุคคล ศูนย์บริการสุขภาพ คณะเภสัชศาสตร์ บริการด้านสุขภาพร่างกาย เครื่องแบบนักศึกษา / บัตรนักศึกษาอิเล็กทรอนิกส์',
    (SELECT type_id FROM news_type WHERE type_name = 'ข่าวประชาสัมพันธ์' LIMIT 1),
    'https://drive.google.com/file/d/1FnJRketlluku27HQqs-mVPnxMJ5wFxqZ/view?usp=sharing&fbclid=IwZXh0bgNhZW0CMTAAAR3M6TbS4tN1DUSa-z2NaM1ekAOELZp7TnsIhJWC6g_dvfz-sD_b0La0S7U_aem_K5_Zhi2eLKhIpJxaeszOlQ'),
('เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล',
    'เปิดรับสมัครแล้ว! ภาควิชาคอมพิวเตอร์ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร เปิดรับสมัครนักศึกษาระดับบัณฑิตศึกษา ปริญญาโท - ปริญญาเอก หลักสูตร IT สำหรับคนทำงาน สู่ผู้เชี่ยวชาญด้านวิจัยและนวัตกรรม เน้นเรียนรู้กระบวนทำวิจัยเพื่อแก้ปัญหาจริง มีทุนผู้ช่วยวิจัยและทุนนำเสนอผลงานในงานประชุมวิชาการ มีทั้งภาคปกติ (เรียนวันธรรมดา) และโครงการพิเศษ (เรียน ส อา) มี 3 แผนการเรียนให้เลือก Data Science, Project Management, DevOps เปิดรับสมัครรอบ 1 วันที่ 16 ธ.ค. 67 - 21 ก.พ. 68 เปิดรับสมัครรอบ 2 วันี่ต่ 3 มี.ค. 68 - 7 พ.ค. 68 สมัครง่าย ๆ ผ่านช่องทางออนไลน์ได้ที่ https://graduate.su.ac.th มีข้อสงสัยสอบถามเพิ่มเติมได้ที่ เพจ Facebook : https://www.facebook.com/computingsu/ หรือ โทร 034-272-923',
    (SELECT type_id FROM news_type WHERE type_name = 'ข่าวประชาสัมพันธ์' LIMIT 1),
    'https://graduate.su.ac.th'),
('กำหนดการรับสมัครและสัมภาษณ์ทุนการศึกษา ประจำปีการศึกษา 2568 (ครั้งที่ 1)',
    'กำหนดการรับสมัครและสัมภาษณ์ทุนการศึกษา ของคณะวิทยาศาสตร์ เพื่อขอทุนการศึกษาสำหรับนักศึกษาปริญญาตรี ภาคการศึกษาต้น ปีการศึกษา 2568 สังกัดคณะวิทย์ ขาดแคลนทุน เกรดไม่ต่ำกว่า 2.00 ขอและยื่นใบสัมครได้ที่ " งานบริการการศึกษา ชั้น 1 อาคารวิทย์ 1 "',
    (SELECT type_id FROM news_type WHERE type_name = 'ทุนการศึกษา' LIMIT 1),
    'https://dsa.su.ac.th/ksu/wp-content/uploads/2025/06/%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B8%81%E0%B8%B2%E0%B8%A8%E0%B8%81%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B4%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%99%E0%B8%B1%E0%B8%81%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2-%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87-%E0%B8%81%E0%B8%B3%E0%B8%AB%E0%B8%99%E0%B8%94%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A3.pdf'),
('ขอแสดงความยินดีกับ ผู้ช่วยศาสตราจารย์ ดร.อรวรรณ เชาวลิต อาจารย์ประจำภาควิชาคอมพิวเตอร์',
    'คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร ขอแสดงความยินดีกับ ผู้ช่วยศาสตราจารย์ ดร.อรวรรณ เชาวลิต ภาควิชาคอมพิวเตอร์ ในโอกาสที่ตีพิมพ์ผลงานวิจัยในวารสาร ICIC Express Letters, Part B: Applications ในฐานข้อมูล Scopus (ScimagoJR, Quartile 4) เรื่อง Multivariate time series forecasting Thailand soybean meal price with deep learning models ICIC Express Letters, Part B: Applications 2025, 16(3), 317-323. http://www.icicelb.org/ellb/contents/2025/3/elb-16-03-10.pdf SDG 8 DECENT WORK AND ECONOMIC GROWTH',
    (SELECT type_id FROM news_type WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'http://www.icicelb.org/ellb/contents/2025/3/elb-16-03-10.pdf'),
('ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์ ดร.ปัญญนัท อ้นพงษ์ เนื่องในโอกาสได้รับการแต่งตั้งให้ดำรงตำแหน่งทางวิชาการ',
    'ภาควิชาคอมพิวเตอร์ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์ ดร.ปัญญนัท อ้นพงษ์ เนื่องในโอกาสได้รับการแต่งตั้งให้ดำรงตำแหน่งทางวิชาการ ตำแหน่ง “ผู้ช่วยศาสตราจารย์” ในสาขาวิชาเทคโนโลยีสารสนเทศ (1806)',
    (SELECT type_id FROM news_type WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'https://www.facebook.com/ScienceSilpakornUniversity/posts/คณะวิทยาศาสตร์-มหาวิทยาลัยศิลปากร-ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์-ดรปัญญนัท/982826210555894/'),
('ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล',
    'ภาควิชาคอมพิวเตอร์ ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล ที่ได้รับรางวัล Best Presentation ในงาน The 6th Asia Joint Conference on Computing (AJCC 2025) ณ เมือง Osaka ประเทศญี่ปุ่นจาก Paper หัวข้อ "Intention Classification of Chinese Topics on Thai Facebook Pages Using Transformer Models with Emotional Features" อาจารย์ที่ปรึกษา อ.ดร.สัจจาภรณ์ ไวจรรยา และ ผศ.ดร.ณัฐโชติ พรหมฤทธิ์',
    (SELECT type_id FROM news_type WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'https://www.facebook.com/computingsu/posts/-ภาควิชาคอมพิวเตอร์-ขอแสดงความยินดีกับ-นายภากร-กัทชลี-นักศึกษา-ปริญญาเอก-หลักสูต/122207080106177331/'),
('ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร',
    'ภาควิชาคอมพิวเตอร์ ขอแสดงความยินดีกับนักศึกษาและอาจารย์ที่ปรึกษา ที่เข้าร่วมนำเสนอผลงานและได้รับรางวัลในงาน การประชุมวิชาการระดับปริญญาตรีด้านคอมพิวเตอร์ภูมิภาคเอเชีย ครั้งที่ 13 (The 13th Asia Undergraduate Conference on Computing: AUCC 2025) ในระหว่างวันที่ 19 - 21 กุมภาพันธ์ พ.ศ. 2568 ณ มหาวิทยาลัยราชภัฏอุตรดิตถ์ จังหวัดอุตรดิตถ์ รางวัลรองชนะเลิศอันดับ 1 Track นวัตกรรม กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร (Development of examination Management system for the Faculty of Science)" ซึ่งผลงานนี้ได้มีการนำไปใช้จริงทดแทนระบบเดิม ในการจัดสอบของคณะฯ เมื่อช่วงสอบกลางภาคการศึกษาที่ผ่านมา นักศึกษาที่เข้านำเสนอ - นาย ณัฐพิสิษ กังวานธรรมกุล - นาย นัทธพงศ์ เป็กทอง - นาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ อาจารย์ที่ปรึกษา ผศ.ดร.ณัฐโชติ พรหมฤทธิ์ และ อ.ดร.สัจจาภรณ์ ไวจรรยา',
    (SELECT type_id FROM news_type WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'https://www.facebook.com/computingsu/posts/-ภาควิชาคอมพิวเตอร์-ขอแสดงความยินดีกับนักศึกษาและอาจารย์ที่ปรึกษา-ที่เข้าร่วมนำเ/122194936622177331/'),
('งานสานสัมพันธ์ภาคคอมพิวเตอร์',
    'เตรียมตัวให้พร้อม! งานสานสัมพันธ์ภาคคอมพิวเตอร์กำลังจะเริ่มขึ้น! เจอกันวันที่ 28 กุมภาพันธ์ 2568 ณ ลานจอดรถตึก 4 เวลา 17:30 - 21:30 น. ธีม: ย้อนยุค ไม่จำกัดช่วงเวลา! จะเป็นยุคหิน มนุษย์ถ้ำ อารยธรรมโบราณ หรือยุคใดก็ได้ จัดเต็มมาให้สุด! มีรางวัลแต่งกายยอดเยี่ยม 3 รางวัล!ใครแต่งตัวเข้าธีมและโดดเด่นที่สุด มีสิทธิ์คว้ารางวัลไปเลย! สนุกไปกับการแสดงสุดพิเศษ, เกมสุดมันส์, ลุ้นรางวัล และดนตรีสดปิดท้าย! ใครลงชื่อไว้แล้ว เจอกันแน่นอน! เตรียมชุดให้พร้อม แล้วมาสนุกไปด้วยกัน!',
    (SELECT type_id FROM news_type WHERE type_name = 'กิจกรรมของภาควิชา' LIMIT 1),
    '');

INSERT INTO news_image(news_id,file_image) VALUES
((SELECT news_id FROM news WHERE title = 'คู่มือแนะนำนักศึกษาใหม่ ปีการศึกษา 2568' LIMIT 1),
    'https://i.postimg.cc/d1p9Dsf1/manual.jpg'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://i.postimg.cc/0jR4606h/degree1.jpg'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://i.postimg.cc/zXkQBWqP/degree2.jpg'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://i.postimg.cc/nLNR5G1M/degree3.jpg'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://i.postimg.cc/gcVSSRLk/degree4.jpg'),
((SELECT news_id FROM news WHERE title = 'กำหนดการรับสมัครและสัมภาษณ์ทุนการศึกษา ประจำปีการศึกษา 2568 (ครั้งที่ 1)' LIMIT 1),
    'https://i.postimg.cc/sD2YYZpL/scholarship.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ ผู้ช่วยศาสตราจารย์ ดร.อรวรรณ เชาวลิต อาจารย์ประจำภาควิชาคอมพิวเตอร์' LIMIT 1),
    'https://i.postimg.cc/YqkWRLQ7/reward-orawan.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์ ดร.ปัญญนัท อ้นพงษ์ เนื่องในโอกาสได้รับการแต่งตั้งให้ดำรงตำแหน่งทางวิชาการ' LIMIT 1),
    'https://i.postimg.cc/T31hNVwz/reward-panyanut.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://i.postimg.cc/gkk0rVvx/reward-phakon1.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://i.postimg.cc/Dzqhx0rV/reward-phakon2.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://i.postimg.cc/DZnTHh6S/reward2nd1.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://i.postimg.cc/7YpjjDYq/reward2nd2.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://i.postimg.cc/1t5dD3Mr/reward2nd3.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://i.postimg.cc/cLy5XmXH/reward2nd4.jpg');

