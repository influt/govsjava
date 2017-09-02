import java.io.*;

public class HelloThreads {
	private final static String FILE_PATH = "./resources/readme.txt";
	private final static int BUF_SIZE = 10;
	private final static Object LOCK_OBJ = new Object();
	private static boolean done = false;

	public static void main(String[] args) {
		ByteArrayOutputStream outStream = new ByteArrayOutputStream();
		
		new Thread(){
			public void run(){
				InputStream fileStream = null;
				int bytesRead = 0;
				byte[] buf = new byte[BUF_SIZE];
				try{
					fileStream = new FileInputStream(FILE_PATH);
					while ((bytesRead=fileStream.read(buf))!=-1){
						outStream.write(buf, 0, bytesRead);
						synchronized (LOCK_OBJ) {
							LOCK_OBJ.wait();
						}
					}
					done = true;
				} catch(InterruptedException e){
					e.printStackTrace();
				} catch(IOException e){
					e.printStackTrace();
				} finally {
					try {
						if (fileStream != null)
							fileStream.close();
						outStream.flush();
					}catch(IOException e){
						e.printStackTrace();
					}
				}
			}
		}.start();
		
		
		try {
			int offset = 0, len, newBytes;
			do {
				byte[] buf = outStream.toByteArray();
				len = buf.length;
				newBytes = len-offset;
				System.out.print(new String(buf, offset, newBytes));
				offset += newBytes;
				synchronized (LOCK_OBJ) {
					LOCK_OBJ.notify();
				}
			} while (!done || newBytes > 0);
			outStream.close();
		}catch(IOException e){
			e.printStackTrace();
		}
	}
}
