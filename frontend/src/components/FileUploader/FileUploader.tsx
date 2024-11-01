import { message, Upload, UploadFile } from 'antd'
import React, { useState } from 'react'
import { InboxOutlined } from '@ant-design/icons'
import { completeMultipart, getMultiUploadUri, getSuccessChunks, newMultiUpload } from '@/api/fileApi'
import { OkKey, TheFileContinueUploadKey, TheFileErrorUploadKey, TheFileIsUploadKey } from '@/locales/locale'
import { useIntl } from 'react-intl'


const { Dragger } = Upload;


type FileUploaderProps = {
    FileType : number,
}


export const FileUploader:React.FC<FileUploaderProps>  = ({FileType})=>{
    const [fileList, setFileList] = useState<UploadFile[]>([]);
    const [uploading, setUploading] = useState(false);
    const CHUNK_SIZE = 5 * 1024 * 1024;
    const intl = useIntl();
    const CONCURRENT_LIMIT = 5;

    // 计算文件的 SHA-256 哈希值
    const calculateSHA256 = (file: File): Promise<string> => {
        return new Promise((resolve, reject) => {
            const reader = new FileReader();
            reader.onload = async (event) => {
                const arrayBuffer = event.target?.result as ArrayBuffer;
                try {
                    const hashBuffer = await crypto.subtle.digest("SHA-256", arrayBuffer);
                    const hashArray = Array.from(new Uint8Array(hashBuffer));
                    const hashHex = hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
                    resolve(hashHex);
                } catch (error) {
                    reject(error);
                }
            };
            reader.onerror = (error) => reject(error);
            reader.readAsArrayBuffer(file);
        });
    };

    // 计算文件的分块数
    const calculateFileChunks = (file: File): number => {
        const fileSize = file.size; // 获取文件大小（字节）
        const chunkCount = Math.ceil(fileSize / CHUNK_SIZE);
        return chunkCount;
    };


    // 上传文件分块
    const uploadMinio = async (url: string, chunk: Blob) => {
        var xhr = new XMLHttpRequest();
        xhr.open('PUT', url, false);
        xhr.setRequestHeader('Content-Type', 'text/plain')
        xhr.send(chunk);
    };

    // 上传文件
    const uploadFile = async (options:any) => {
        console.log(options);
        const {onSuccess,onError,file,onProgress }= options;
        const fileHash = await calculateSHA256(file);
        const totalChunks = calculateFileChunks(file);
        file.status = 'uploading';
        setFileList([...fileList, file]);


        let isUpload:Boolean = false
        let isRecord:boolean = false
        let chunksID:string = ""
        const needUploadChunksID:number[] = [];
        let  progressCnt = 0;


        const uploadSingleChunk = async (chunkIndex: number) => {
            const start = chunkIndex * CHUNK_SIZE;
            const end = Math.min(start + CHUNK_SIZE, file.size);
            const chunk = file.slice(start, end);
            const uploadUrl = await getUploadUri(fileHash, chunkIndex+1, end-start);
            await uploadMinio(uploadUrl, chunk);
            console.log(`Uploaded chunk ${chunkIndex + 1} of ${totalChunks}`);
        };

        const uploadChunksConcurrently = async () => {
            const uniqueChunkIDs = Array.from(new Set(needUploadChunksID)); // 去重

            let currentIndex = 0;

            while (currentIndex < uniqueChunkIDs.length) {
                const currentBatch = uniqueChunkIDs.slice(currentIndex, currentIndex + CONCURRENT_LIMIT);

                const batch = currentBatch.map(async (chunkIndex) => {
                    await uploadSingleChunk(chunkIndex); // 上传单个块
                    progressCnt += 1;
                    file.percent = ((progressCnt / totalChunks) * 100>95?95:(progressCnt / totalChunks) * 100).toFixed(2);
                    setFileList([...fileList, file]);
                });

                await Promise.all(batch);
                currentIndex += CONCURRENT_LIMIT;
            }
        };

        // 判断是否已经上传
        await getSuccessChunks({
            file_hash: fileHash,
        }).then((res) => {
            isUpload = res.data.is_upload;
            chunksID = res.data.chunks;
            isRecord = res.data.is_record;
        });

        if(isUpload){
            file.status = 'done';
            setFileList([...fileList, file]);
            message.info(intl.formatMessage({ id: TheFileIsUploadKey }));
            return;
        }else{
            if (!isRecord){
                await newMultiUpload({
                    file_hash: fileHash,
                    file_type: FileType,
                    file_name: file.name,
                    file_size: file.size,
                    chunk_total_number: totalChunks,
                }).catch((error) => {
                    file.status = 'error';
                    setFileList([...fileList, file]);
                    message.error(intl.formatMessage({ id: TheFileErrorUploadKey }));
                })


            const chunkIDArraySS = chunksID ? chunksID.split(",") : []; // 如果 chunksID 为空，则返回空数组
            const chunkIDArray = chunkIDArraySS.map((item) => parseInt(item, 10)).filter(Number.isInteger); // 过滤掉非整数值
            progressCnt = chunkIDArray.length;

            if(chunkIDArray.length !== totalChunks){
                message.info(intl.formatMessage({ id: TheFileContinueUploadKey }));
                for (let i = 0; i < totalChunks; i++) {
                    if (!chunkIDArray.includes(i)) {
                        needUploadChunksID.push(i);
                    }
                }
            }

                for(let i = 0; i < totalChunks; i++){
                    needUploadChunksID.push(i);
                }
            }

        }

        try {
            await uploadChunksConcurrently();
        } catch (error) {
            file.status = 'error';
            setFileList([...fileList, file]);
            message.error(intl.formatMessage({ id: TheFileErrorUploadKey }));
        }

        try {
            await completeMultipart(
                {
                    file_hash: fileHash,
                }
            )
            message.success(intl.formatMessage({ id: OkKey }))
        }catch (error) {
            file.status = 'error';
            setFileList([...fileList, file]);
            message.error(intl.formatMessage({ id: TheFileErrorUploadKey }));
        }finally {
            setUploading(false);
            file.status = 'done';
            file.percent = 100;
            setFileList([...fileList, file]);
        }




    };


    const getUploadUri = async (fileHash:string,chunk_id: number,chunk_size:number)=>{
        return await getMultiUploadUri({
            file_hash: fileHash,
            chunk_id: chunk_id,
            chunk_size: chunk_size,
        }).then((res) => {
            return res.data.uri;
        });
    }

    return (
        <>
            <Dragger
                beforeUpload={async (file) => {

                }}
                onRemove={(file) => {
                    setFileList((prevList) => prevList.filter((item) => item.uid !== file.uid));
                }}
                fileList={fileList}
                multiple={true}
                customRequest={async (options) => {
                    await uploadFile(options);
                }
                }
            >
                <p className="ant-upload-drag-icon">
                    <InboxOutlined />
                </p>
                <p className="ant-upload-text">Click or drag file to this area to upload</p>
                <p className="ant-upload-hint">
                    Support for a single or bulk upload. Strictly prohibited from uploading company data or other
                    banned files.
                </p>
            </Dragger>

        </>
    )
}